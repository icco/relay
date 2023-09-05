// Code generated by prana; DO NOT EDIT.

// Package model contains an object model of database schema 'public'
// Auto-generated at Tue, 05 Sep 2023 14:52:14 UTC

package model

import "time"

// GeographyColumn represents a data base table 'geography_columns'
type GeographyColumn struct {
	// FTableCatalog represents a database column 'f_table_catalog' of type 'NAME NULL'
	FTableCatalog *string `db:"f_table_catalog,null"`
	// FTableSchema represents a database column 'f_table_schema' of type 'NAME NULL'
	FTableSchema *string `db:"f_table_schema,null"`
	// FTableName represents a database column 'f_table_name' of type 'NAME NULL'
	FTableName *string `db:"f_table_name,null"`
	// FGeographyColumn represents a database column 'f_geography_column' of type 'NAME NULL'
	FGeographyColumn *string `db:"f_geography_column,null"`
	// CoordDimension represents a database column 'coord_dimension' of type 'INTEGER(32) NULL'
	CoordDimension *int `db:"coord_dimension,null"`
	// Srid represents a database column 'srid' of type 'INTEGER(32) NULL'
	Srid *int `db:"srid,null"`
	// Type represents a database column 'type' of type 'TEXT NULL'
	Type *string `db:"type,null"`
}

// GeometryColumn represents a data base table 'geometry_columns'
type GeometryColumn struct {
	// FTableCatalog represents a database column 'f_table_catalog' of type 'CHARACTER VARYING(256) NULL'
	FTableCatalog *string `db:"f_table_catalog,null,size=256"`
	// FTableSchema represents a database column 'f_table_schema' of type 'NAME NULL'
	FTableSchema *string `db:"f_table_schema,null"`
	// FTableName represents a database column 'f_table_name' of type 'NAME NULL'
	FTableName *string `db:"f_table_name,null"`
	// FGeometryColumn represents a database column 'f_geometry_column' of type 'NAME NULL'
	FGeometryColumn *string `db:"f_geometry_column,null"`
	// CoordDimension represents a database column 'coord_dimension' of type 'INTEGER(32) NULL'
	CoordDimension *int `db:"coord_dimension,null"`
	// Srid represents a database column 'srid' of type 'INTEGER(32) NULL'
	Srid *int `db:"srid,null"`
	// Type represents a database column 'type' of type 'CHARACTER VARYING(30) NULL'
	Type *string `db:"type,null,size=30"`
}

// Lock represents a data base table 'locks'
type Lock struct {
	// Name represents a database column 'name' of type 'CHARACTER VARYING(255) PRIMARY KEY NOT NULL'
	Name string `db:"name,primary_key,not_null,size=255"`
	// RecordVersionNumber represents a database column 'record_version_number' of type 'BIGINT(64) NULL'
	RecordVersionNumber *int64 `db:"record_version_number,null"`
	// Data represents a database column 'data' of type 'BYTEA NULL'
	Data []byte `db:"data,null"`
	// Owner represents a database column 'owner' of type 'CHARACTER VARYING(255) NULL'
	Owner *string `db:"owner,null,size=255"`
}

// Message represents a data base table 'messages'
type Message struct {
	// ID represents a database column 'id' of type 'INTEGER(32) PRIMARY KEY NOT NULL'
	ID int `db:"id,primary_key,not_null"`
	// Message represents a database column 'message' of type 'JSONB NOT NULL'
	Message []byte `db:"message,not_null"`
	// Published represents a database column 'published' of type 'BOOLEAN NOT NULL'
	Published bool `db:"published,not_null"`
	// CreatedAt represents a database column 'created_at' of type 'TIMESTAMP WITH TIME ZONE NOT NULL'
	CreatedAt time.Time `db:"created_at,not_null"`
	// UpdatedAt represents a database column 'updated_at' of type 'TIMESTAMP WITH TIME ZONE NOT NULL'
	UpdatedAt time.Time `db:"updated_at,not_null"`
	// DeletedAt represents a database column 'deleted_at' of type 'TIMESTAMP WITH TIME ZONE NULL'
	DeletedAt *time.Time `db:"deleted_at,null"`
	// Content represents a database column 'content' of type 'TEXT NOT NULL'
	Content string `db:"content,not_null"`
}

// SpatialRefSy represents a data base table 'spatial_ref_sys'
type SpatialRefSy struct {
	// Srid represents a database column 'srid' of type 'INTEGER(32) PRIMARY KEY NOT NULL'
	Srid int `db:"srid,primary_key,not_null"`
	// AuthName represents a database column 'auth_name' of type 'CHARACTER VARYING(256) NULL'
	AuthName *string `db:"auth_name,null,size=256"`
	// AuthSrid represents a database column 'auth_srid' of type 'INTEGER(32) NULL'
	AuthSrid *int `db:"auth_srid,null"`
	// Srtext represents a database column 'srtext' of type 'CHARACTER VARYING(2048) NULL'
	Srtext *string `db:"srtext,null,size=2048"`
	// Proj4text represents a database column 'proj4text' of type 'CHARACTER VARYING(2048) NULL'
	Proj4text *string `db:"proj4text,null,size=2048"`
}
