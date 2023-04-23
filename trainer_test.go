package main

import (
	"math"
	"testing"
	"trainder-api/models"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

func TestFilterTrainerWithLimit(t *testing.T) {
	var specialties []string

	result, err := models.FindFilteredTrainer(specialties, 1, 0, math.MaxInt)

	assert.Equal(t, nil, err)

	assert.Equal(t, 1, len(result))

}

func TestFilterTrainerWithNoFeeLimit(t *testing.T) {
	var specialties []string

	result, err := models.FindFilteredTrainer(specialties, 1, 0, 0)

	assert.Equal(t, nil, err)

	assert.Equal(t, 1, len(result))

}

func TestFilterTrainerWithOneSpecialty(t *testing.T) {
	var specialties []string
	specialties = append(specialties, "Weight Loss")
	specialties = append(specialties, "Rehabilitation")

	result, err := models.FindFilteredTrainer(specialties, 1, 0, 0)

	assert.Equal(t, nil, err)

	assert.LessOrEqual(t, 1, len(result))

	if len(result) > 0 {
		done := false
		for _, spec := range result[0].TrainerInfo.Specialty {
			if slices.Contains(specialties, spec) {
				done = true
				break
			}
		}
		assert.True(t, done)
	}
}

func TestFilterTrainerWithManySpecialty(t *testing.T) {
	var specialties []string
	specialties = append(specialties, "Weight Loss")

	result, err := models.FindFilteredTrainer(specialties, 1, 0, 0)

	assert.Equal(t, nil, err)

	assert.LessOrEqual(t, 1, len(result))

	if len(result) > 0 {
		done := false
		for _, spec := range result[0].TrainerInfo.Specialty {
			if slices.Contains(specialties, spec) {
				done = true
				break
			}
		}
		assert.True(t, done)
	}
}

func TestFilterTrainerWithInvalidSpecialty(t *testing.T) {
	var specialties []string
	specialties = append(specialties, "Bla Bla")

	_, err := models.FindFilteredTrainer(specialties, 1, 0, 0)

	assert.Equal(t, nil, err)

}
