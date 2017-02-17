/**
 * Created by tuanha-01 on 5/13/2016.
 */
"use strict";

define(['angularAMD', 'jquery', 'ajaxService', 'alertsService', 'myApp.autoComplete', 'profileQuestionnairesService'], function(angularAMD, $) {
    var injectParams = ['$scope', '$rootScope', '$state', '$window', 'moment', '$uibModalInstance', 'alertsService', 'profileQuestionnairesService', '$stateParams', 'Constants', 'profileQuestionnaireLineEditRatings'];

    var ratingsMaintenanceController = function($scope, $rootScope, $state, $window, moment, $uibModalInstance, alertsService, profileQuestionnairesService, $stateParams, Constants, profileQuestionnaireLineEditRatings) {

        $scope.ProfileQuestionnaireLineEditRatings = profileQuestionnaireLineEditRatings;
        $scope.ProfileQuestionnaireLines = [];
        $scope.ProfileQuestionnaires = [];
        $scope.Constants = Constants;

        $scope.initializeController = function() {
            $scope.selectProfileQuestionnaireHeaderID = "";

            $scope.getProfileQuestionnaires();
        };

        $scope.getProfileQuestionnaires = function() {
            var profileQuestionnaireInquiry = new Object();

            profileQuestionnaireInquiry.Search = "";
            profileQuestionnaireInquiry.SortExpression = "Code";
            profileQuestionnaireInquiry.SortDirection = "ASC";
            profileQuestionnaireInquiry.FetchSize = ""; //get all

            profileQuestionnairesService.getProfileQuestionnaires(profileQuestionnaireInquiry, $scope.profileQuestionnairesInquiryCompleted, $scope.profileQuestionnairesInquiryError);
        };

        $scope.profileQuestionnairesInquiryCompleted = function(response, status) {
            $scope.ProfileQuestionnaires = response.Data.ProfileQuestionnaires;
        };

        $scope.profileQuestionnairesInquiryError = function(response, status) {
            alertsService.RenderErrorMessage(response.Error);
        };

        $scope.selectProfileQuestionnaire = function(_profileQuestionnaireHeaderID) {
            $scope.selectProfileQuestionnaireHeaderID = _profileQuestionnaireHeaderID;

            $scope.getProfileQuestionnaireLines(_profileQuestionnaireHeaderID);
        };

        $scope.getProfileQuestionnaireLines = function(_profileQuestionnaireHeaderID) {
            var __profileQuestionnaire = new Object();
            __profileQuestionnaire.HeaderID = _profileQuestionnaireHeaderID
            profileQuestionnairesService.getProfileQuestionnaireLines(__profileQuestionnaire, $scope.getProfileQuestionnaireLinesCompleted, $scope.getProfileQuestionnaireLinesError);
        };

        $scope.getProfileQuestionnaireLinesCompleted = function(response, status) {
            alertsService.RenderSuccessMessage(response.ReturnMessage);

            var profileQuestionnaireLines = response.Data.ProfileQuestionnaireLines;
            var profileQuestionnaireLineDeletes = [];
            var found_question = false;
            for (var i = 0, len = profileQuestionnaireLines.length; i < len; i++) {
                profileQuestionnaireLines[i].RecCreated = new moment.unix(profileQuestionnaireLines[i].RecCreated).toDate();
                profileQuestionnaireLines[i].RecModified = new moment.unix(profileQuestionnaireLines[i].RecModified).toDate();

                if (profileQuestionnaireLines[i].ProfileQuestionnaireHeaderID == $scope.ProfileQuestionnaireLineEditRatings.ProfileQuestionnaireHeaderID) {
                    if (profileQuestionnaireLines[i].LineNo == $scope.ProfileQuestionnaireLineEditRatings.LineNo) {
                        found_question = true;
                        profileQuestionnaireLineDeletes.push(profileQuestionnaireLines[i])
                    } else if (found_question && profileQuestionnaireLines[i].Type == Constants.ProfileQuestionaireLineTypes[1].Code) {
                        profileQuestionnaireLineDeletes.push(profileQuestionnaireLines[i])
                    } else if (found_question && profileQuestionnaireLines[i].Type == Constants.ProfileQuestionaireLineTypes[0].Code) {
                        found_question = false;
                    }
                }
            }
            profileQuestionnaireLines = profileQuestionnaireLines.filter(function(el, index, array) {
                return !profileQuestionnaireLineDeletes.find(function(_el) {
                    return _el.ID === el.ID
                })
            })
            $scope.ProfileQuestionnaireLines = profileQuestionnaireLines;
        };

        $scope.getProfileQuestionnaireLinesError = function(response, status) {
            alertsService.RenderErrorMessage(response.Error);
        }

        $scope.validationOptions = {
            rules: {
                "Point[]": {
                    number: true
                }
            }
        };

        $scope.ok = function(form) {
            if (form.validate()) {
                $uibModalInstance.close($scope.ProfileQuestionnaireLineEditRatings);
            }
        };

        $scope.cancel = function(form) {
            $uibModalInstance.dismiss('cancel');
        };

        $scope.changeProfileQuestionaire = function() {

        };
    };

    ratingsMaintenanceController.$inject = injectParams;
    angularAMD.controller('RatingsMaintenanceController', ratingsMaintenanceController);
});