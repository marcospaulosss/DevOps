package repositories

type PreferenceQueryBuilder struct{}

func NewPreferenceQueryBuilder() PreferenceQueryBuilder {
	return PreferenceQueryBuilder{}
}

func (this PreferenceQueryBuilder) ReadOne() string {
	return `SELECT type, content FROM preferences 
	WHERE type = $1 AND content IS NOT NULL 
	ORDER BY created_at DESC LIMIT 1;`
}

func (this PreferenceQueryBuilder) Update() string {
	return `INSERT INTO preferences (type, content)
		VALUES (:type, :content)
		ON CONFLICT (type)
		DO UPDATE SET content = :content WHERE preferences.type = :type
		RETURNING type, content;`
}
