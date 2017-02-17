/**
 * Created by tuanha-01 on 5/13/2016.
 */
"use strict";

define(['angularAMD', 'jquery', 'ajaxService', 'alertsService', 'myApp.autoComplete', 'profileQuestionnairesService'], function(angularAMD, $) {
    var injectParams = ['$scope', '$rootScope', '$state', '$window', 'moment', '$uibModalInstance', 'alertsService', 'profileQuestionnairesService', '$stateParams', 'Constants', 'profileQuestionnaireLine'];

    var ratingsMaintenanceController = function($scope, $rootScope, $state, $window, moment, $uibModalInstance, alertsService, profileQuestionnairesService, $stateParams, Constants, profileQuestionnaireLine) {

        $scope.ProfileQuestionnaireLine = profileQuestionnaireLine;
        $scope.ProfileQuestionnaireLines = [];
        $scope.ProfileQuestionnaires = [];

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
            for (var i = 0, len = profileQuestionnaireLines.length; i < len; i++) {
                profileQuestionnaireLines[i].RecCreated = new moment.unix(profileQuestionnaireLines[i].RecCreated).toDate();
                profileQuestionnaireLines[i].RecModified = new moment.unix(profileQuestionnaireLines[i].RecModified).toDate();
            }

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
                $uibModalInstance.close($scope.ProfileQuestionnaireLine);
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