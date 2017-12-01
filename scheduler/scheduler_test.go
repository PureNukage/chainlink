package scheduler_test

import (
	. "github.com/onsi/gomega"
	"github.com/smartcontractkit/chainlink-go/internal/cltest"
	"github.com/smartcontractkit/chainlink-go/models"
	"github.com/smartcontractkit/chainlink-go/scheduler"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadingSavedSchedules(t *testing.T) {
	RegisterTestingT(t)
	store := cltest.Store()
	defer store.Close()

	j := models.NewJob()
	j.Schedule = models.Schedule{Cron: "* * * * *"}
	jobWoCron := models.NewJob()
	_ = store.Save(&j)
	_ = store.Save(&jobWoCron)

	sched := scheduler.New(store.ORM)
	err := sched.Start()
	assert.Nil(t, err)
	defer sched.Stop()

	jobRuns := []models.JobRun{}
	Eventually(func() []models.JobRun {
		_ = store.Where("JobID", j.ID, &jobRuns)
		return jobRuns
	}).Should(HaveLen(1))

	err = store.Where("JobID", jobWoCron.ID, &jobRuns)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(jobRuns), "No jobs should be created without the scheduler")
}

func TestAddJob(t *testing.T) {
	RegisterTestingT(t)
	store := cltest.Store()
	defer store.Close()

	sched := scheduler.New(store.ORM)
	_ = sched.Start()
	defer sched.Stop()

	j := models.NewJob()
	j.Schedule = models.Schedule{Cron: "* * * * *"}
	_ = store.Save(&j)
	sched.AddJob(j)

	jobRuns := []models.JobRun{}
	Eventually(func() []models.JobRun {
		_ = store.Where("JobID", j.ID, &jobRuns)
		return jobRuns
	}).Should(HaveLen(1))
}
