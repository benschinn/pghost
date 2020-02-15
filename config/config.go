package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	SourceConnection               string `yaml:"source_connection"`
	SourceConnectionForReplication string `yaml:"source_connection_for_replication"`
	SourceTableName                string `yaml:"source_table_name"`
	SourceSchemaName               string `yaml:"source_schema_name"`

	DestinationConnection string `yaml:"destination_connection"`
	DestinationTableName  string `yaml:"destination_table_name"`
	DestinationSchemaName string `yaml:"destination_schema_name"`

	PublicationName            string `yaml:"publication_name"`
	ReplicationSlotName        string `yaml:"replication_slot_name"`
	ReplicationSlotIsTemporary bool   `yaml:"replication_slot_is_temporary"`

	CopyBatchSize           int  `yaml:"copy_batch_size"`
	CopyWorkerCount         int  `yaml:"copy_worker_count"`
	CopyUseKeysetPagination bool `yaml:"copy_use_keyset_pagination"`
}

func (c *Config) generateMissingValues() error {
	if len(c.ReplicationSlotName) == 0 {
		c.ReplicationSlotName = "pghost_created_replication_slot"
	}
	return nil
}

func ParseConfig(path string) (*Config, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	c := new(Config)
	err = yaml.Unmarshal(bytes, c)
	if err != nil {
		return nil, err
	}
	err = c.generateMissingValues()
	return c, err
}