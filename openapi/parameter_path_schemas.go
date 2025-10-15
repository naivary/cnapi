package openapi

var StringParameter = &Schema{
	Type: StringType,
}

var IntegerParameter = &Schema{
	Type: IntegerType,
}

var UUIDParameter = &Schema{
	Type:   StringType,
	Format: "uuid",
}

func RegExpParam(pattern string) *Schema {
	return &Schema{
		Type:    StringType,
		Pattern: pattern,
	}
}
