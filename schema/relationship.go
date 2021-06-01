package schema

// RelationshipType relationship type
type RelationshipType string

const (
	HasOne    RelationshipType = "has_one"      // HasOneRel has one relationship
	HasMany   RelationshipType = "has_many"     // HasManyRel has many relationship
	BelongsTo RelationshipType = "belongs_to"   // BelongsToRel belongs to relationship
	Many2Many RelationshipType = "many_to_many" // Many2ManyRel many to many relationship
	has       RelationshipType = "has"
)
