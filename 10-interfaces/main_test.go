package main

import "testing"

func TestMain(t *testing.T) {
	t.Run("processTruck", func(t *testing.T) {
		t.Run("should load and unload a truck cargo", func(t *testing.T) {
			truck := &NormalTruck{id: "T1"}

			// Act
			err := processTruck(truck)

			// Assert
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
	})

}
