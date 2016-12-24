package v2_test

import (
	"errors"

	"code.cloudfoundry.org/cli/command"
	"code.cloudfoundry.org/cli/command/commandfakes"
	"code.cloudfoundry.org/cli/command/v2"
	"code.cloudfoundry.org/cli/util/configv3"
	"code.cloudfoundry.org/cli/util/ui"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
)

var _ = FDescribe("target Command", func() {
	var (
		cmd    v2.TargetCommand
		testUI *ui.UI
		// fakeActor  *v2fakes.FakeDeleteOrganizationActor
		fakeConfig *commandfakes.FakeConfig
		input      *Buffer
		executeErr error
	)

	BeforeEach(func() {
		input = NewBuffer()
		testUI = ui.NewTestUI(input, NewBuffer(), NewBuffer())
		// fakeActor = new(v2fakes.FakeDeleteOrganizationActor)
		fakeConfig = new(commandfakes.FakeConfig)

		cmd = v2.TargetCommand{
			UI: testUI,
			// Actor:  fakeActor,
			Config: fakeConfig,
		}
		fakeConfig.BinaryNameReturns("faceman")

		// cmd.RequiredArgs.Organization = "some-org"
	})

	JustBeforeEach(func() {
		executeErr = cmd.Execute(nil)
	})

	Context("when the user is not logged in", func() {
		It("returns an error", func() {
			Expect(executeErr).To(MatchError(command.NotLoggedInError{
				BinaryName: "faceman",
			}))
		})
	})

	Context("when config.CurrentUser returns an error", func() {
		var someError error
		BeforeEach(func() {
			someError = errors.New("some-current-user-error")
			fakeConfig.AccessTokenReturns("some-access-token")
			fakeConfig.CurrentUserReturns(configv3.User{}, someError)
		})

		It("returns the error", func() {
			Expect(executeErr).To(MatchError(someError))
		})
	})

	Context("when the user is logged in", func() {
		BeforeEach(func() {
			fakeConfig.TargetReturns("some-api-target")
			fakeConfig.APIVersionReturns("1.2.3")
			fakeConfig.AccessTokenReturns("some-access-token")
			fakeConfig.RefreshTokenReturns("some-refresh-token")
			fakeConfig.CurrentUserReturns(configv3.User{
				Name: "some-user",
			}, nil)
		})

		Context("when no arguments are given", func() {
			Context("when no org or space is targetted", func() {
				BeforeEach(func() {
					fakeConfig.TargetedOrganizationReturns(configv3.Organization{})
					fakeConfig.TargetedSpaceReturns(configv3.Space{})
				})

				It("displays no org or space targetted", func() {
					Expect(executeErr).ToNot(HaveOccurred())

					Expect(testUI.Out).To(Say("API endpoint:   some-api-target \\(API version: 1.2.3\\)"))
					Expect(testUI.Out).To(Say("User:           some-user"))
					Expect(testUI.Out).To(Say("No org or space targeted, use 'faceman target -o ORG -s SPACE'"))
				})
			})
		})
	})

	Context("when an org argument is given", func() {
		BeforeEach(func() {
			// cmd.Organization = "some-org"
		})

		Context("when getting the org returns an error", func() {
			It("returns the error", func() {
			})
		})

		Context("when getting the org does not return an error", func() {
			It("sets the org", func() {
			})

			It("tells the user that the org has been targetted", func() {
			})
		})
	})

	// When -o
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//

	// 		Context("when the delete actor org returns an error", func() {
	// 			Context("when the organization does not exist", func() {
	// 				BeforeEach(func() {
	// 					fakeActor.DeleteOrganizationReturns(
	// 						v2action.Warnings{"warning-1", "warning-2"},
	// 						v2action.OrganizationNotFoundError{
	// 							Name: "some-org",
	// 						},
	// 					)
	// 				})

	// 				It("returns an OrganizationNotFoundError and displays all warnings", func() {
	// 					Expect(executeErr).NotTo(HaveOccurred())

	// 					Expect(testUI.Out).To(Say("Deleting org some-org as some-user..."))

	// 					Expect(fakeActor.DeleteOrganizationCallCount()).To(Equal(1))
	// 					orgName := fakeActor.DeleteOrganizationArgsForCall(0)
	// 					Expect(orgName).To(Equal("some-org"))

	// 					Expect(testUI.Err).To(Say("warning-1"))
	// 					Expect(testUI.Err).To(Say("warning-2"))

	// 					Expect(testUI.Out).To(Say("Org some-org does not exist."))
	// 					Expect(testUI.Out).To(Say("OK"))
	// 				})
	// 			})
	// 		})
	// 	})

	// 	// Testing the prompt.
	// 	Context("when the '-f' flag is not provided", func() {
	// 		Context("when the user chooses the default", func() {
	// 			BeforeEach(func() {
	// 				input.Write([]byte("\n"))
	// 			})

	// 			It("does not delete the org", func() {
	// 				Expect(executeErr).ToNot(HaveOccurred())

	// 				Expect(testUI.Out).To(Say("Delete cancelled"))

	// 				Expect(fakeActor.DeleteOrganizationCallCount()).To(Equal(0))
	// 			})
	// 		})

	// 		Context("when the user inputs no", func() {
	// 			BeforeEach(func() {
	// 				input.Write([]byte("n\n"))
	// 			})

	// 			It("does not delete the org", func() {
	// 				Expect(executeErr).ToNot(HaveOccurred())

	// 				Expect(testUI.Out).To(Say("Delete cancelled"))

	// 				Expect(fakeActor.DeleteOrganizationCallCount()).To(Equal(0))
	// 			})
	// 		})

	// 		Context("when the user inputs yes", func() {
	// 			BeforeEach(func() {
	// 				input.Write([]byte("y\n"))
	// 			})

	// 			It("deletes the org", func() {
	// 				Expect(executeErr).ToNot(HaveOccurred())

	// 				Expect(testUI.Out).To(Say("Really delete the org some-org and everything associated with it\\?>> \\[yN\\]:"))
	// 				Expect(testUI.Out).To(Say("Deleting org some-org as some-user..."))

	// 				Expect(fakeActor.DeleteOrganizationCallCount()).To(Equal(1))
	// 				orgName := fakeActor.DeleteOrganizationArgsForCall(0)
	// 				Expect(orgName).To(Equal("some-org"))

	// 				Expect(testUI.Out).To(Say("OK"))
	// 			})
	// 		})

	// 		Context("when the user input is invalid", func() {
	// 			BeforeEach(func() {
	// 				input.Write([]byte("e\n\n"))
	// 			})

	// 			It("asks the user again", func() {
	// 				Expect(executeErr).NotTo(HaveOccurred())

	// 				Expect(testUI.Out).To(Say("Really delete the org some-org and everything associated with it\\?>> \\[yN\\]:"))
	// 				Expect(testUI.Out).To(Say("invalid input \\(not y, n, yes, or no\\)"))
	// 				Expect(testUI.Out).To(Say("Really delete the org some-org and everything associated with it\\?>> \\[yN\\]:"))

	// 				Expect(fakeActor.DeleteOrganizationCallCount()).To(Equal(0))
	// 			})
	// 		})

	// 		Context("when displaying the prompt returns an error", func() {
	// 			// if nothing is written to input, display bool prompt returns EOF
	// 			It("returns the error", func() {
	// 				Expect(executeErr).To(MatchError("EOF"))
	// 			})
	// 		})
	// 	})
	// })
})
