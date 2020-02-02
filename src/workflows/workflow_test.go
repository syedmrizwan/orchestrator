package workflows

import (
	"errors"
	"github.com/stretchr/testify/suite"
	"go.uber.org/cadence/testsuite"
	"github.com/syedmrizwan/orchestrator/src/activities"
	"testing"
)

type UnitTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite

	env *testsuite.TestWorkflowEnvironment
}

func (s *UnitTestSuite) SetupTest() {
	s.env = s.NewTestWorkflowEnvironment()
}

func (s *UnitTestSuite) AfterTest(suiteName, testName string) {
	s.env.AssertExpectations(s.T())
}

func (s *UnitTestSuite) Test_DemoWorkflow() {
	s.env.ExecuteWorkflow(DemoWorkflow)

	s.True(s.env.IsWorkflowCompleted())
	s.NoError(s.env.GetWorkflowError())
}

func (s *UnitTestSuite) Test_DemoWorkflow_Error() {
	s.env.OnActivity(activities.GetNameActivity).Return("", errors.New("oops")).Once()
	s.env.OnActivity(activities.GetNameActivity).Return("Mock", nil).Once()
	s.env.ExecuteWorkflow(DemoWorkflow)
	// s.True(s.env.IsWorkflowCompleted())

	// s.NotNil(s.env.GetWorkflowError())
	// s.Equal("SimpleActivityFailure", s.env.GetWorkflowError().Error())

	s.True(s.env.IsWorkflowCompleted())
	s.NoError(s.env.GetWorkflowError())
	s.env.AssertExpectations(s.T())
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, new(UnitTestSuite))
}
