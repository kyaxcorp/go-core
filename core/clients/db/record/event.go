package record

func (r *Record) callOnBeforeDelete() {
	if _model, ok := r.modelStruct.(interface{ RecordBeforeDelete(r *Record) }); ok {
		_model.RecordBeforeDelete(r)
	}

	// TODO: call other methods from record...
}

func (r *Record) callOnError() {
	if _model, ok := r.modelStruct.(interface{ RecordError(r *Record) }); ok {
		_model.RecordError(r)
	}

	// TODO: call other methods from record...
}

func (r *Record) callOnDBError() {
	if _model, ok := r.modelStruct.(interface{ RecordDBError(r *Record) }); ok {
		_model.RecordDBError(r)
	}

	// TODO: call other methods from record...
}

func (r *Record) callOnSaveError() {
	if _model, ok := r.modelStruct.(interface{ RecordSaveError(r *Record) }); ok {
		_model.RecordSaveError(r)
	}

	// TODO: call other methods from record...
}

func (r *Record) callOnDeleteError() {
	if _model, ok := r.modelStruct.(interface{ RecordDeleteError(r *Record) }); ok {
		_model.RecordDeleteError(r)
	}

	// TODO: call other methods from record...
}

func (r *Record) callOnAfterDelete() {
	if _model, ok := r.modelStruct.(interface{ RecordAfterDelete(r *Record) }); ok {
		_model.RecordAfterDelete(r)
	}

	// TODO: call other methods from record...
}

func (r *Record) callOnBeforeForceDelete() {
	if _model, ok := r.modelStruct.(interface{ RecordBeforeForceDelete(r *Record) }); ok {
		_model.RecordBeforeForceDelete(r)
	}

	// TODO: call other methods from record...
}

func (r *Record) callOnAfterForceDelete() {
	if _model, ok := r.modelStruct.(interface{ RecordAfterForceDelete(r *Record) }); ok {
		_model.RecordAfterForceDelete(r)
	}

	// TODO: call other methods from record...
}

func (r *Record) callOnBeforeSave() {
	if _model, ok := r.modelStruct.(interface{ RecordBeforeSave(r *Record) }); ok {
		_model.RecordBeforeSave(r)
	}

	// TODO: call other methods from record...
}

func (r *Record) callOnAfterDbDataLoaded() {
	if _model, ok := r.modelStruct.(interface{ RecordAfterDBDataLoaded(r *Record) }); ok {
		_model.RecordAfterDBDataLoaded(r)
	}

	// TODO: call other methods from record...
}

func (r *Record) callOnAfterSave() {
	if _model, ok := r.modelStruct.(interface{ RecordAfterSave(r *Record) }); ok {
		_model.RecordAfterSave(r)
	}

	// TODO: call other methods from record...
}

func (r *Record) callOnBeforeInsert() {
	if _model, ok := r.modelStruct.(interface{ RecordBeforeInsert(r *Record) }); ok {
		_model.RecordBeforeInsert(r)
	}

	// TODO: call other methods from record...
}

func (r *Record) callOnAfterInsert() {
	if _model, ok := r.modelStruct.(interface{ RecordAfterInsert(r *Record) }); ok {
		_model.RecordAfterInsert(r)
	}

	// TODO: call other methods from record...
}

func (r *Record) callOnBeforeUpdate() {
	if _model, ok := r.modelStruct.(interface{ RecordBeforeUpdate(r *Record) }); ok {
		_model.RecordBeforeUpdate(r)
	}

	// TODO: call other methods from record...
}

func (r *Record) callOnAfterUpdate() {
	if _model, ok := r.modelStruct.(interface{ RecordAfterUpdate(r *Record) }); ok {
		_model.RecordAfterUpdate(r)
	}

	// TODO: call other methods from record...
}
