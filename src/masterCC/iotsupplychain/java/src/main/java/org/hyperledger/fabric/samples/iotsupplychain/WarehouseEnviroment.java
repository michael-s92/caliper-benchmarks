package org.hyperledger.fabric.samples.iotsupplychain;

import java.util.Objects;

import org.hyperledger.fabric.contract.annotation.DataType;
import org.hyperledger.fabric.contract.annotation.Property;

import com.owlike.genson.annotation.JsonProperty;

@DataType()
public class WarehouseEnviroment extends BatchAnalys {

	@Property()
	private final Integer moisturePercent;
	
	public Integer getMoisturePercent() {
		
		return moisturePercent;
	}
	
	public WarehouseEnviroment(@JsonProperty("moisture_percent") final Integer moisturePercent) {
		
		this.moisturePercent = moisturePercent;
	}
	
	@Override
	public boolean equals(Object obj) {
		
		if(this == obj) {
			return true;
		}
		
		if((obj == null) || (getClass() != obj.getClass())) {
			return false;
		}
		
		WarehouseEnviroment other = (WarehouseEnviroment) obj;
		return Objects.deepEquals(getMoisturePercent(), other.getMoisturePercent());
	}
	
	@Override
	public int hashCode() {
		
		return Objects.hash(getMoisturePercent());
	}
	
	@Override
	public String toString() {

		return "WarehouseEnviroment( " + moisturePercent + "%  moisture )";
	}
}
