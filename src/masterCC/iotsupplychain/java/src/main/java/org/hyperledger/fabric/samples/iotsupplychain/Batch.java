package org.hyperledger.fabric.samples.iotsupplychain;


import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

import org.hyperledger.fabric.contract.annotation.DataType;
import org.hyperledger.fabric.contract.annotation.Property;

import com.owlike.genson.annotation.JsonProperty;

@DataType()
public class Batch {

	private static String docType = "batch-doc";

	@Property()
	private final String id;

	@Property()
	private final String farmerId;

	@Property()
	private final String warehouseId;

	@Property()
	private Boolean inWarehouse;

	@Property()
	private LabResult labResult;

	@Property()
	private final List<WarehouseEnviroment> warehouseAnalysis;

	@Property()
	private String discount;

	public String getId() {
		return id;
	}

	public String getFarmerId() {
		return farmerId;
	}

	public String getWarehouseId() {
		return warehouseId;
	}

	public Boolean getInWarehouse() {
		return inWarehouse;
	}

	public LabResult getLabResult() {
		return labResult;
	}

	public List<WarehouseEnviroment> getWarehouseAnalysis(){
		return warehouseAnalysis;
	}

	public String getDiscount() {
		return discount;
	}

	public Batch(@JsonProperty("id") final String id, @JsonProperty("farmerId") final String farmerId, @JsonProperty("warehouseId") final String warehouseId,
			@JsonProperty("inWarehouse") final Boolean inWarehouse, @JsonProperty("labResult") final LabResult labResult,
			@JsonProperty("warehouseAnalysis") final List<WarehouseEnviroment> warehouseAnalysis, 
			@JsonProperty("discount") final String discount) {

		this.id = id;
		this.farmerId = farmerId;
		this.warehouseId = warehouseId;
		if( inWarehouse == null ) {
			this.inWarehouse = Boolean.TRUE;
		}else {
			this.inWarehouse = inWarehouse;
		}
		this.labResult = labResult;
		if( warehouseAnalysis == null ) {
			this.warehouseAnalysis = new ArrayList<WarehouseEnviroment>();
		}else {
			this.warehouseAnalysis = warehouseAnalysis;
		}
		this.discount = discount;
	}

	public Batch(final String id, final String farmerId, final String warehouseId) {
		this.id = id;
		this.farmerId = farmerId;
		this.warehouseId = warehouseId;
		this.inWarehouse = Boolean.TRUE;
		this.warehouseAnalysis = new ArrayList<WarehouseEnviroment>();
	}

	public void leaveWarehouse() {
		inWarehouse = Boolean.FALSE;
	}

	public void calculateDiscount() {

		double avgMoisture = 0;
		for (WarehouseEnviroment e : warehouseAnalysis) {
			avgMoisture += e.getMoisturePercent().doubleValue();
		}
		avgMoisture = avgMoisture / this.warehouseAnalysis.size();

		double discount = 0;

		if (avgMoisture > 12) {
			discount += (avgMoisture - 12) * 4;
		}
		if (this.labResult.getImpurityPercent() > 3) {
			discount += (this.labResult.getImpurityPercent() - 3) * 2.5;
		}
		if (this.labResult.getBrokenPercent() > 5) {
			discount += (this.labResult.getBrokenPercent() - 5) * 1;
		}
		if (this.labResult.getDamagedPercent() > 3) {
			discount += (this.labResult.getDamagedPercent() - 3) * 3.5;
		}

		this.discount = discount + "";
	}

	public void pushWarehouseAnalysis(final Integer moisturePercent) {

		this.warehouseAnalysis.add(new WarehouseEnviroment(moisturePercent));
	}

	public void storeLabResults(final String labId, final Integer impurity_percent,
			final Integer broken_percent, final Integer damaged_percent, final Integer greenisch_percent) {
		this.labResult = new LabResult(labId, impurity_percent, broken_percent, damaged_percent, greenisch_percent);
	}

	@Override
	public boolean equals(Object obj) {
		
		if(obj == this) return true;
		if((obj != null) && (obj.getClass() != getClass())) return false;
		
		Batch other = (Batch) obj;
		return Objects.deepEquals(new String[] { getId(), getFarmerId(), getWarehouseId(), getDiscount() }, 
				new String[] { other.getId(), other.getFarmerId(), other.getWarehouseId(), other.getDiscount() }) 
				&& Objects.deepEquals(getInWarehouse(), other.getInWarehouse()) && Objects.deepEquals(getWarehouseAnalysis(), other.getWarehouseAnalysis())
				&& Objects.deepEquals(getLabResult(), other.getLabResult());
	}

	@Override
	public int hashCode() {

		return Objects.hash(getId(), getFarmerId(), getWarehouseId(), getInWarehouse(), getWarehouseAnalysis(), getLabResult(), getDiscount());
	}

	@Override
	public String toString() {

		return "Batch[ " + getId() + " ]( " + getFarmerId() + ", " + getWarehouseId() + " )";
	}
}
