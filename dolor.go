	task := &datastore.Entity{
		Key: datastore.NameKey("Task", "sampleTask", nil),
		Properties: []datastore.Property{
			datastore.Property{
				Name:  "Category",
				Value: "Personal",
			},
			datastore.Property{
				Name:  "Done",
				Value: false,
			},
			datastore.Property{
				Name:  "Priority",
				Value: 4,
			},
			datastore.Property{
				Name:  "PercentComplete",
				Value: 10.0,
			},
			datastore.Property{
				Name:  "Created",
				Value: time.Now(),
			},
		},
	}
	_, err := client.Put(ctx, task.Key, task)  
