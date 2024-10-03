package builder

type VehicleBuilder struct {
	v VehicleProduct
}

func (v *VehicleBuilder) SetWheels() BuildProcess {
	v.v.Wheels = 2
	return v
}

func (v *VehicleBuilder) SetSeats() BuildProcess {
	v.v.Seats = 2
	return v
}

func (v *VehicleBuilder) SetStructure() BuildProcess {
	v.v.Structure = "Bike"
	return v
}

func (v *VehicleBuilder) GetVehicle() VehicleProduct {
	return v.v
}
