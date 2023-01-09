// Copyright Axis Communications AB.
//
// For a full list of individual contributors, please see the commit history.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// THIS FILE IS AUTOMATICALLY GENERATED AND MUST NOT BE EDITED BY HAND.

package eiffelevents

import "reflect"

type majorEventVersion struct {
	structType    reflect.Type
	latestVersion string
}

// eventTypeTable maps the major versions of each event to a struct containing
// a type reference to the Go type used to represent that event and the most
// recent version of that event within that major version.
var eventTypeTable = map[string]map[int64]majorEventVersion{
	"EiffelActivityCanceledEvent": {
		1: majorEventVersion{reflect.TypeOf(ActivityCanceledV1{}), "1.1.0"},
		2: majorEventVersion{reflect.TypeOf(ActivityCanceledV2{}), "2.0.0"},
		3: majorEventVersion{reflect.TypeOf(ActivityCanceledV3{}), "3.2.0"},
	},
	"EiffelActivityFinishedEvent": {
		1: majorEventVersion{reflect.TypeOf(ActivityFinishedV1{}), "1.1.0"},
		2: majorEventVersion{reflect.TypeOf(ActivityFinishedV2{}), "2.0.0"},
		3: majorEventVersion{reflect.TypeOf(ActivityFinishedV3{}), "3.3.0"},
	},
	"EiffelActivityStartedEvent": {
		1: majorEventVersion{reflect.TypeOf(ActivityStartedV1{}), "1.1.0"},
		2: majorEventVersion{reflect.TypeOf(ActivityStartedV2{}), "2.0.0"},
		3: majorEventVersion{reflect.TypeOf(ActivityStartedV3{}), "3.0.0"},
		4: majorEventVersion{reflect.TypeOf(ActivityStartedV4{}), "4.3.0"},
	},
	"EiffelActivityTriggeredEvent": {
		1: majorEventVersion{reflect.TypeOf(ActivityTriggeredV1{}), "1.1.0"},
		2: majorEventVersion{reflect.TypeOf(ActivityTriggeredV2{}), "2.0.0"},
		3: majorEventVersion{reflect.TypeOf(ActivityTriggeredV3{}), "3.0.0"},
		4: majorEventVersion{reflect.TypeOf(ActivityTriggeredV4{}), "4.2.0"},
	},
	"EiffelAnnouncementPublishedEvent": {
		1: majorEventVersion{reflect.TypeOf(AnnouncementPublishedV1{}), "1.1.0"},
		2: majorEventVersion{reflect.TypeOf(AnnouncementPublishedV2{}), "2.0.0"},
		3: majorEventVersion{reflect.TypeOf(AnnouncementPublishedV3{}), "3.2.0"},
	},
	"EiffelArtifactCreatedEvent": {
		1: majorEventVersion{reflect.TypeOf(ArtifactCreatedV1{}), "1.1.0"},
		2: majorEventVersion{reflect.TypeOf(ArtifactCreatedV2{}), "2.0.0"},
		3: majorEventVersion{reflect.TypeOf(ArtifactCreatedV3{}), "3.3.0"},
	},
	"EiffelArtifactPublishedEvent": {
		1: majorEventVersion{reflect.TypeOf(ArtifactPublishedV1{}), "1.1.0"},
		2: majorEventVersion{reflect.TypeOf(ArtifactPublishedV2{}), "2.0.0"},
		3: majorEventVersion{reflect.TypeOf(ArtifactPublishedV3{}), "3.3.0"},
	},
	"EiffelArtifactReusedEvent": {
		1: majorEventVersion{reflect.TypeOf(ArtifactReusedV1{}), "1.1.0"},
		2: majorEventVersion{reflect.TypeOf(ArtifactReusedV2{}), "2.0.0"},
		3: majorEventVersion{reflect.TypeOf(ArtifactReusedV3{}), "3.2.0"},
	},
	"EiffelCompositionDefinedEvent": {
		1: majorEventVersion{reflect.TypeOf(CompositionDefinedV1{}), "1.1.0"},
		2: majorEventVersion{reflect.TypeOf(CompositionDefinedV2{}), "2.0.0"},
		3: majorEventVersion{reflect.TypeOf(CompositionDefinedV3{}), "3.3.0"},
	},
	"EiffelConfidenceLevelModifiedEvent": {
		1: majorEventVersion{reflect.TypeOf(ConfidenceLevelModifiedV1{}), "1.1.0"},
		2: majorEventVersion{reflect.TypeOf(ConfidenceLevelModifiedV2{}), "2.0.0"},
		3: majorEventVersion{reflect.TypeOf(ConfidenceLevelModifiedV3{}), "3.2.0"},
	},
	"EiffelEnvironmentDefinedEvent": {
		1: majorEventVersion{reflect.TypeOf(EnvironmentDefinedV1{}), "1.1.0"},
		2: majorEventVersion{reflect.TypeOf(EnvironmentDefinedV2{}), "2.0.0"},
		3: majorEventVersion{reflect.TypeOf(EnvironmentDefinedV3{}), "3.3.0"},
	},
	"EiffelFlowContextDefinedEvent": {
		1: majorEventVersion{reflect.TypeOf(FlowContextDefinedV1{}), "1.1.0"},
		2: majorEventVersion{reflect.TypeOf(FlowContextDefinedV2{}), "2.0.0"},
		3: majorEventVersion{reflect.TypeOf(FlowContextDefinedV3{}), "3.2.0"},
	},
	"EiffelIssueDefinedEvent": {
		1: majorEventVersion{reflect.TypeOf(IssueDefinedV1{}), "1.0.0"},
		2: majorEventVersion{reflect.TypeOf(IssueDefinedV2{}), "2.0.0"},
		3: majorEventVersion{reflect.TypeOf(IssueDefinedV3{}), "3.2.0"},
	},
	"EiffelIssueVerifiedEvent": {
		1: majorEventVersion{reflect.TypeOf(IssueVerifiedV1{}), "1.1.0"},
		2: majorEventVersion{reflect.TypeOf(IssueVerifiedV2{}), "2.0.0"},
		3: majorEventVersion{reflect.TypeOf(IssueVerifiedV3{}), "3.0.0"},
		4: majorEventVersion{reflect.TypeOf(IssueVerifiedV4{}), "4.2.0"},
	},
	"EiffelSourceChangeCreatedEvent": {
		1: majorEventVersion{reflect.TypeOf(SourceChangeCreatedV1{}), "1.1.0"},
		2: majorEventVersion{reflect.TypeOf(SourceChangeCreatedV2{}), "2.0.0"},
		3: majorEventVersion{reflect.TypeOf(SourceChangeCreatedV3{}), "3.0.0"},
		4: majorEventVersion{reflect.TypeOf(SourceChangeCreatedV4{}), "4.2.0"},
	},
	"EiffelSourceChangeSubmittedEvent": {
		1: majorEventVersion{reflect.TypeOf(SourceChangeSubmittedV1{}), "1.1.0"},
		2: majorEventVersion{reflect.TypeOf(SourceChangeSubmittedV2{}), "2.0.0"},
		3: majorEventVersion{reflect.TypeOf(SourceChangeSubmittedV3{}), "3.2.0"},
	},
	"EiffelTestCaseCanceledEvent": {
		1: majorEventVersion{reflect.TypeOf(TestCaseCanceledV1{}), "1.1.0"},
		2: majorEventVersion{reflect.TypeOf(TestCaseCanceledV2{}), "2.0.0"},
		3: majorEventVersion{reflect.TypeOf(TestCaseCanceledV3{}), "3.2.0"},
	},
	"EiffelTestCaseFinishedEvent": {
		1: majorEventVersion{reflect.TypeOf(TestCaseFinishedV1{}), "1.1.0"},
		2: majorEventVersion{reflect.TypeOf(TestCaseFinishedV2{}), "2.0.0"},
		3: majorEventVersion{reflect.TypeOf(TestCaseFinishedV3{}), "3.3.0"},
	},
	"EiffelTestCaseStartedEvent": {
		1: majorEventVersion{reflect.TypeOf(TestCaseStartedV1{}), "1.1.0"},
		2: majorEventVersion{reflect.TypeOf(TestCaseStartedV2{}), "2.0.0"},
		3: majorEventVersion{reflect.TypeOf(TestCaseStartedV3{}), "3.3.0"},
	},
	"EiffelTestCaseTriggeredEvent": {
		1: majorEventVersion{reflect.TypeOf(TestCaseTriggeredV1{}), "1.1.0"},
		2: majorEventVersion{reflect.TypeOf(TestCaseTriggeredV2{}), "2.0.0"},
		3: majorEventVersion{reflect.TypeOf(TestCaseTriggeredV3{}), "3.2.0"},
	},
	"EiffelTestExecutionRecipeCollectionCreatedEvent": {
		1: majorEventVersion{reflect.TypeOf(TestExecutionRecipeCollectionCreatedV1{}), "1.0.0"},
		2: majorEventVersion{reflect.TypeOf(TestExecutionRecipeCollectionCreatedV2{}), "2.1.0"},
		3: majorEventVersion{reflect.TypeOf(TestExecutionRecipeCollectionCreatedV3{}), "3.0.0"},
		4: majorEventVersion{reflect.TypeOf(TestExecutionRecipeCollectionCreatedV4{}), "4.3.0"},
	},
	"EiffelTestSuiteFinishedEvent": {
		1: majorEventVersion{reflect.TypeOf(TestSuiteFinishedV1{}), "1.1.0"},
		2: majorEventVersion{reflect.TypeOf(TestSuiteFinishedV2{}), "2.0.0"},
		3: majorEventVersion{reflect.TypeOf(TestSuiteFinishedV3{}), "3.3.0"},
	},
	"EiffelTestSuiteStartedEvent": {
		1: majorEventVersion{reflect.TypeOf(TestSuiteStartedV1{}), "1.1.0"},
		2: majorEventVersion{reflect.TypeOf(TestSuiteStartedV2{}), "2.0.0"},
		3: majorEventVersion{reflect.TypeOf(TestSuiteStartedV3{}), "3.3.0"},
	},
}
