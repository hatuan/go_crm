/**
 * Created by tuanha-01 on 5/13/2016.
 */
"use strict";

define(['angularAMD', 'ajaxService', 'alertsService', 'myApp.autoComplete', 'profileQuestionnairesService'], function (angularAMD, $) {
    var injectParams = ['$scope', '$rootScope', '$state', '$window', 'moment', 'alertsService', 'profileQuestionnairesService', '$stateParams', 'Constants'];

    var profileQuestionnaireLinesMaintenanceController = function ($scope, $rootScope, $state, $window, moment, alertsService, profileQuestionnairesService, $stateParams, Constants) {

        $scope.initializeController = function () {
            $rootScope.applicationModule = "ProfileQuestionnaireLinesMaintenance";
            $rootScope.alerts = [];

            $scope.profileQuestionnaireHeaderID = ($stateParams.headerID || "");

            $scope.Constants = Constants;

            $scope.ProfileQuestionnaireLines = [];

            var getProfileQuestionnaireLines = new Object();
            getProfileQuestionnaireLines.HeaderID = $scope.profileQuestionnaireHeaderID
            profileQuestionnairesService.getProfileQuestionnaireLines(getProfileQuestionnaireLines, $scope.getProfileQuestionnaireLinesCompleted, $scope.getProfileQuestionnaireLinesError);
        };

        $scope.getProfileQuestionnaireLinesCompleted = function (response, status) {
            alertsService.RenderSuccessMessage(response.ReturnMessage);

            var profileQuestionnaireLines = response.Data.ProfileQuestionnaireLines;
            for (var i = 0, len = profileQuestionnaireLines.length; i < len; i++) {
                profileQuestionnaireLines[i].RecCreated = new moment.unix(profileQuestionnaireLines[i].RecCreated).toDate();
                profileQuestionnaireLines[i].RecModified = new moment.unix(profileQuestionnaireLines[i].RecModified).toDate();
            }

            $scope.ProfileQuestionnaireLines = profileQuestionnaireLines;
            $scope.TotalRows = response.TotalRows;
        };

        $scope.getProfileQuestionnaireLinesError = function (response, status) {
            alertsService.RenderErrorMessage(response.Error);
        }

        $scope.validationOptions = {
            rules: {
                "Description[]": {
                    required: true
                },
            }
        };

        $scope.update = function (form) {
            if (form.validate({
                errorPlacement: function (error, element) {
                    return true;
                }
            })) {
                var profileQuestionnaire = new Object();
                profileQuestionnaire.HeaderID = $scope.profileQuestionnaireHeaderID

                var profileQuestionnaireLines = $scope.ProfileQuestionnaireLines;
                for (var i = 0, len = profileQuestionnaireLines.length; i < len; i++) {
                    profileQuestionnaireLines[i].RecCreated = new moment(profileQuestionnaireLines[i].RecCreated).unix();
                    profileQuestionnaireLines[i].RecModified = new moment(profileQuestionnaireLines[i].RecModified).unix();
                }

                profileQuestionnairesService.updateProfileQuestionnaireLines(profileQuestionnaire, profileQuestionnaireLines, $scope.profileQuestionnaireLinesUpdateCompleted, $scope.profileQuestionnaireLinesUpdateError)
            }
        };

        $scope.profileQuestionnaireLinesUpdateCompleted = function (response, status) {
            alertsService.RenderSuccessMessage(response.ReturnMessage);
            
            setTimeout(function() {
                $state.go('profileQuestionnaireMaintenance', { ID: $scope.profileQuestionnaireHeaderID });
            }, 1000);
        };

        $scope.profileQuestionnaireLinesUpdateError = function (response, status) {
            alertsService.RenderErrorMessage(response.Error);
        }

        $scope.cancel = function (form) {
            setTimeout(function () {
                $state.go('profileQuestionnaireMaintenance', { ID: $scope.profileQuestionnaireHeaderID });
            }, 10);
        };

        $scope.addLines = function () {
            var profileQuestionnaireLine = $scope.createProfileQuestionnaireLineObject();
            $scope.ProfileQuestionnaireLines.push(profileQuestionnaireLine);
        }

        $scope.lineDetail = function (profileQuestionnaireLine) {

        }

        $scope.createProfileQuestionnaireLineObject = function () {
            var profileQuestionnaireLine = new Object();
            profileQuestionnaireLine.ID = "";
            profileQuestionnaireLine.Type = $scope.Constants.ProfileQuestionaireLineTypes[0].Code;
            profileQuestionnaireLine.ProfileQuestionnaireHeaderID = $scope.profileQuestionnaireHeaderID;
            profileQuestionnaireLine.Description = "";
            profileQuestionnaireLine.Priority = $scope.Constants.Priorities[2].Code;
            profileQuestionnaireLine.MultipleAnswers = $scope.Constants.BooleanTypes[0].Code;
            profileQuestionnaireLine.AutoContactClassification = $scope.Constants.BooleanTypes[0].Code;
            profileQuestionnaireLine.CustomerClassField = 0;
            profileQuestionnaireLine.VendorClassField = 0;
            profileQuestionnaireLine.ContactClassField = 0;
            profileQuestionnaireLine.StartingDateFormula = "";
            profileQuestionnaireLine.EndingDateFormula = "";
            profileQuestionnaireLine.ClassificationMethod = 0;
            profileQuestionnaireLine.SortingMethod = 0;
            profileQuestionnaireLine.FromValue = 0;
            profileQuestionnaireLine.ToValue = 0;

            profileQuestionnaireLine.Status = $scope.Constants.Status[1].Code;
            profileQuestionnaireLine.ClientID = "";
            profileQuestionnaireLine.OrganizationID = "";
            profileQuestionnaireLine.RecCreatedByID = $rootScope.currentUser.ID;
            profileQuestionnaireLine.RecCreatedByUser = $rootScope.currentUser.Name;
            profileQuestionnaireLine.RecCreated = new Date();
            profileQuestionnaireLine.RecModifiedByID = $rootScope.currentUser.ID;
            profileQuestionnaireLine.RecModifiedByUser = $rootScope.currentUser.Name;
            profileQuestionnaireLine.RecModified = new Date();

            return profileQuestionnaireLine;
        }
    };

    profileQuestionnaireLinesMaintenanceController.$inject = injectParams;
    angularAMD.controller('ProfileQuestionnaireLinesMaintenanceController', profileQuestionnaireLinesMaintenanceController);
});
