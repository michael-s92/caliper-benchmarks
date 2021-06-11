package org.hyperledger.fabric.samples.eureka;

import org.hyperledger.fabric.contract.Context;
import org.hyperledger.fabric.contract.ContractInterface;
import org.hyperledger.fabric.contract.annotation.Contract;
import org.hyperledger.fabric.contract.annotation.Default;
import org.hyperledger.fabric.contract.annotation.Transaction;

import com.owlike.genson.Genson;

@Contract(name = "eureka")
@Default
public class EurekaContract implements ContractInterface{

	private final Genson genson = new Genson();
	
	@Transaction()
	public void init(final Context ctx) {
		
	}
	
	@Transaction()
	public void doNothing(final Context ctx) {
		
	}
	
	@Transaction()
	public void initLedger(final Context ctx) {
		
	}
}
