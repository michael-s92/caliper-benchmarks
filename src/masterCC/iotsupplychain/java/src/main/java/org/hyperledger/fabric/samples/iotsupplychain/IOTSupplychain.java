package org.hyperledger.fabric.samples.iotsupplychain;

import org.hyperledger.fabric.contract.Context;
import org.hyperledger.fabric.contract.annotation.Contract;
import org.hyperledger.fabric.contract.annotation.Default;
import org.hyperledger.fabric.contract.annotation.Transaction;
import org.hyperledger.fabric.shim.ChaincodeException;
import org.hyperledger.fabric.shim.ChaincodeStub;
import org.json.JSONObject;

import com.owlike.genson.Genson;

@Contract(name = "iotsupplychain")
@Default
public class IOTSupplychain {

	private final Genson genson = new Genson();
	
	@Transaction()
	public void init(final Context ctx){
		
	}
	
	@Transaction()
	public void doNothing(final Context ctx){
		
	}
	
	@Transaction()
	public void initLedger(final Context ctx){
		/*
		 
		  for(const batchJson of seeds.initBatchs){
           
            let batch = Batch.fromJSON(batchJson);
            await ctx.stub.putState(batchJson.id, Buffer.from(JSON.stringify(batch)));
          }

		 */
	}
	
	@Transaction()
	public void sendBanchToWarehouse(final Context ctx, final String batchId, final String farmerId, final String warehouseId){
		
		if(batchId == null || batchId.isEmpty()) {
			throw new ChaincodeException(CCErrors.KEY_NULL_OR_EMPTY.toString());
		}
		
		if(farmerId == null || farmerId.isEmpty()) {
			throw new ChaincodeException(CCErrors.KEY_NULL_OR_EMPTY.toString());
		}
		
		if(warehouseId == null || warehouseId.isEmpty()) {
			throw new ChaincodeException(CCErrors.KEY_NULL_OR_EMPTY.toString());
		}
		
		ChaincodeStub stub = ctx.getStub();
        // check if that key already existes
		String batchState = stub.getStringState(batchId);
		if(!batchState.isEmpty()) {
			throw new ChaincodeException(CCErrors.STATE_ALREADY_EXISTS.toString());
		}
		
		Batch batch = new Batch(batchId, farmerId, warehouseId);
		batchState = genson.serialize(batch);
		stub.putStringState(batchId, batchState);
	}
	
	@Transaction()
	public void saveChemicalAnalysToBatch(final Context ctx, final String batchId, final String labId, 
			final String impurity_percent, final String broken_percent, final String damaged_percent, final String greenisch_percent) {
		
		if(batchId == null || batchId.isEmpty()) {
			throw new ChaincodeException(CCErrors.KEY_NULL_OR_EMPTY.toString());
		}
		
		if(labId == null || labId.isEmpty()) {
			throw new ChaincodeException(CCErrors.KEY_NULL_OR_EMPTY.toString());
		}
		
		if(impurity_percent == null || impurity_percent.isEmpty()) {
			throw new ChaincodeException(CCErrors.KEY_NULL_OR_EMPTY.toString());
		} 
		Integer impurity = Integer.parseInt(impurity_percent);
		if(impurity < 0 || impurity > 100) {
			throw new ChaincodeException(CCErrors.PERCENT_OUT_OF_RANGE.toString());
		}
		
		if(broken_percent == null || broken_percent.isEmpty()) {
			throw new ChaincodeException(CCErrors.KEY_NULL_OR_EMPTY.toString());
		}
		Integer broken = Integer.parseInt(impurity_percent);
		if(broken < 0 || broken > 100) {
			throw new ChaincodeException(CCErrors.PERCENT_OUT_OF_RANGE.toString());
		}
		
		if(damaged_percent == null || damaged_percent.isEmpty()) {
			throw new ChaincodeException(CCErrors.KEY_NULL_OR_EMPTY.toString());
		}
		Integer damaged = Integer.parseInt(impurity_percent);
		if(damaged < 0 || damaged > 100) {
			throw new ChaincodeException(CCErrors.PERCENT_OUT_OF_RANGE.toString());
		}
		
		if(greenisch_percent == null || greenisch_percent.isEmpty()) {
			throw new ChaincodeException(CCErrors.KEY_NULL_OR_EMPTY.toString());
		}
		Integer greenisch = Integer.parseInt(impurity_percent);
		if(greenisch < 0 || greenisch > 100) {
			throw new ChaincodeException(CCErrors.PERCENT_OUT_OF_RANGE.toString());
		}
		
		ChaincodeStub stub = ctx.getStub();
		String batchState = stub.getStringState(batchId);
		if(batchState.isEmpty()) throw new ChaincodeException(CCErrors.STATE_DOES_NOT_EXIST.toString());
		
		Batch batch = genson.deserialize(batchState, Batch.class);
		
		if( batch.getLabResult() != null ) throw new ChaincodeException(CCErrors.ATTRIBUT_ALREADY_POPULATED.toString());
		
		batch.storeLabResults(labId, impurity, broken, damaged, greenisch);
		
		batchState = genson.serialize(batch);
		stub.putStringState(batchId, batchState);
	}
	
	@Transaction()
	public void savePsysicalAnalysToBatch(final Context ctx, final String warehouseId, final String moisture_percent) {
		
		if(warehouseId == null || warehouseId.isEmpty()) {
			throw new ChaincodeException(CCErrors.KEY_NULL_OR_EMPTY.toString());
		}
		
		if(moisture_percent == null || moisture_percent.isEmpty()) {
			throw new ChaincodeException(CCErrors.KEY_NULL_OR_EMPTY.toString());
		}
		Integer moisture = Integer.parseInt(moisture_percent);
		if(moisture < 0 || moisture > 100) {
			throw new ChaincodeException(CCErrors.PERCENT_OUT_OF_RANGE.toString());
		}
		
		ChaincodeStub stub = ctx.getStub();
		
		JSONObject query = new JSONObject();
		JSONObject inner = new JSONObject();
		inner.put("warehouseId", warehouseId);
		inner.put("inWarehouse", true);
		query.put("selector", inner);
		stub.getQueryResult(query.toString());
		
		/*
		  // fetch all batchs that are in warehouse
        let batchQueryString = {};
        batchQueryString.selector = {
            warehouseId: warehouseId,
            inWarehouse: true
        };

        let resultIterator = await ctx.stub.getQueryResult(JSON.stringify(batchQueryString));
        let results = await Helper.getAllResultsKeyValue(resultIterator);

        // store enviroment data to them
        for (const obj of results) {
            let tmpBatch = obj.record;
            tmpBatch.pushWarehouseAnalysis(moisture);
            await ctx.stub.putState(obj.key, Buffer.from(JSON.stringify(tmpBatch)));
        }
		 */
	}
	
	@Transaction()
	public void moveBatchToFoodCompany(final Context ctx, final String batchId) {
		
		/*
		 if (batchId.length <= 0) {
            throw new Error("BatchId must be non empty string");
        }

        // fetch batch
        let batchAsBytes = await ctx.stub.getState(batchId);
        if (!Helper.objExists) {
            throw new Error(`Batch ${batchId} doesnt exist`);
        }

        let batchJson = {};
        try {
            batchJson = JSON.parse(batchAsBytes.toString());
        } catch (err) {
            throw new Error(`Failed to parse batch ${batchId}, err: ${err}`);
        }

        let batch = Batch.fromJSON(batchJson);

        // check if there is no LabResults
        if (batch.inWarehouse === undefined || batch.inWarehouse === false) {
            throw new Error(`Batch ${batchId} already has moved`);
        }

        // move batch to producer
        batch.leaveWarehous();

        // save batch to state
        await ctx.stub.putState(batchId, Buffer.from(JSON.stringify(batch)));
		 */
	}
	
	@Transaction()
	public void calculateDiscount(final Context ctx, final String batchId) {
		/*
		 if (batchId.length <= 0) {
            throw new Error("BatchId must be non empty string");
        }

        // fetch batch
        let batchAsBytes = await ctx.stub.getState(batchId);
        if (!Helper.objExists) {
            throw new Error(`Batch ${batchId} doesnt exist`);
        }

        let batchJson = {};
        try {
            batchJson = JSON.parse(batchAsBytes.toString());
        } catch (err) {
            throw new Error(`Failed to parse batch ${batchId}, err: ${err}`);
        }

        let batch = Batch.fromJSON(batchJson);

        // move batch to producer
        batch.calculateDiscount();

        //emit event about discount
        let payload = {
            batch_id: batchId,
            discount: batch.discount
        };
        ctx.stub.setEvent('discount_calculated_event', Buffer.from(JSON.stringify(payload)));

        // save batch to state
        await ctx.stub.putState(batchId, Buffer.from(JSON.stringify(batch)));
		 */
	}
	
	@Transaction()
	public void getFarmerRating(final Context ctx, final String farmerId) {
		/*
		 if (farmerId.length <= 0) {
            throw new Error("farmerId must be non empty string");
        }

        // fetch all batchs for farmer
        let batchQueryString = {};
        batchQueryString.selector = {
            farmerId: farmerId,
            inWarehouse: false
        };

        let resultIterator = await ctx.stub.getQueryResult(JSON.stringify(batchQueryString));
        let results = await Helper.getAllResults(resultIterator);

        let raiting = 0;

        // store enviroment data to them
        for (const batch of results) {

            batch.calculateDiscount();
            raiting += (1 - batch.discount);
        }

        raiting = raiting / results.length;

        let payload = {
            farmerId: farmerId,
            raiting: raiting,
            number: results.length
        };

        return Buffer.from(JSON.stringify(payload));
		 */
	}
}