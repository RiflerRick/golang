package main

type DataBlock struct {
	dataPointer *string
}

/*
for computing less than for the actual data
*/
func (d *DataBlock) LessThan(dParam DataBlock) bool {
	dResolved := *(d.dataPointer)
	dParamResolved := *(dParam.dataPointer)
	if dResolved < dParamResolved {
		return true
	}
	return false
}

/*
for computing more than for actual data
*/
func (d *DataBlock) MoreThan(dParam DataBlock) bool {
	dResolved := *(d.dataPointer)
	dParamResolved := *(dParam.dataPointer)
	if dResolved > dParamResolved {
		return true
	}
	return false
}

/*
for computing equals for actual data
*/
func (d *DataBlock) Equals(dParam DataBlock) bool {
	dResolved := *(d.dataPointer)
	dParamResolved := *(dParam.dataPointer)
	if dResolved == dParamResolved {
		return true
	}
	return false
}
