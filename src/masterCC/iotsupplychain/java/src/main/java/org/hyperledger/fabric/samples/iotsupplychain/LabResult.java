package org.hyperledger.fabric.samples.iotsupplychain;

import java.util.Objects;

import org.hyperledger.fabric.contract.annotation.DataType;
import org.hyperledger.fabric.contract.annotation.Property;

import com.owlike.genson.annotation.JsonProperty;

@DataType()
public class LabResult extends BatchAnalys{

	@Property()
	private final String labId;
	
	@Property()
	private final Integer impurityPercent;
	
	@Property()
	private final Integer brokenPercent;
	
	@Property()
	private final Integer damagedPercent;
	
	@Property()
	private final Integer greenischPercent;
	
	public String getLabId() {
		return labId;
	}
	
	public Integer getImpurityPercent() {
		return impurityPercent;
	}
	
	public Integer getBrokenPercent() {
		return brokenPercent;
	}
	
	public Integer getDamagedPercent() {
		return damagedPercent;
	}
	
	public Integer getGreenischPercent() {
		return greenischPercent;
	}
	
	public LabResult(@JsonProperty("labId") final String labId, @JsonProperty("impurity_percent") final Integer impurityPercent, 
			@JsonProperty("broken_percent") final Integer brokenPercent, @JsonProperty("damaged_percent") final Integer damagedPercent, 
			@JsonProperty("greenisch_percent") final Integer greenischPercent) {
		
		this.labId = labId;
		this.impurityPercent = impurityPercent;
		this.brokenPercent = brokenPercent;
		this.damagedPercent = damagedPercent;
		this.greenischPercent = greenischPercent;
	}
	
	@Override
	public boolean equals(Object obj) {

		if(this == obj) {
			return true;
		}
		
		if((obj == null) || obj.getClass() != getClass()) {
			return false;
		}
		
		LabResult other = (LabResult) obj;
		
		return Objects.deepEquals(getLabId(), other.getLabId()) &&
				Objects.deepEquals(new Integer[] { getImpurityPercent(), getBrokenPercent(), getDamagedPercent(), getGreenischPercent() }, 
				new Integer[] { other.getImpurityPercent(), other.getBrokenPercent(), other.getDamagedPercent(), other.getGreenischPercent()});
	}
	
	@Override
	public int hashCode() {

		return Objects.hash(getLabId(), getImpurityPercent(), getBrokenPercent(), getDamagedPercent(), getGreenischPercent());
	}
	
	@Override
	public String toString() {

		return "LabResults[ " + getLabId() + " ]( " + getImpurityPercent() + ", " + getBrokenPercent() + ", " + getDamagedPercent() + ", " + getGreenischPercent() + " )";
	}
}
