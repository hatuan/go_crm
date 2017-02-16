/**
 * Created by tuanha-01 on 5/13/2016.
 */
"use strict";

define(['angularAMD', 'jquery', 'ajaxService', 'alertsService', 'myApp.autoComplete', 'profileQuestionnairesService'], function(angularAMD, $) {
    var injectParams = ['$scope', '$rootScope', '$state', '$window', 'moment', '$uibModalInstance', 'alertsService', 'profileQuestionnairesService', '$stateParams', 'Constants', 'profileQuestionnaireLine'];

    var ratingsMaintenanceController = function($scope, $rootScope, $state, $window, moment, $uibModalInstance, alertsService, profileQuestionnairesService, $stateParams, Constants, profileQuestionnaireLine) {

        $scope.ProfileQuestionnaireLine = profileQuestionnaireLine;
        $scope.ProfileQuestionnaireLines = [];

        $scope.initializeController = function() {
            $scope.ProfileQuestionnaireHeaderID = $scope.ProfileQuestionnaireLine.HeaderID;
            $scope.ProfileQuestionnaireHeaderCode = "";

            var dataURL = new Object();
            dataURL.HeaderID = $scope.ProfileQuestionnaireHeaderID;
            profileQuestionnairesService.getProfileQuestionnaireLinesAndRatings(dataURL, $scope.getProfileQuestionnaireLinesAndRatingsCompleted, $scope.getProfileQuestionnaireLinesAndRatingsError);
        };

        $scope.getProfileQuestionnaireLinesAndRatingsCompleted = function(response, status) {
            alertsService.RenderSuccessMessage(response.ReturnMessage);

            var profileQuestionnaire = response.Data.ProfileQuestionnaire
            var profileQuestionnaireLines = response.Data.ProfileQuestionnaireLines;

            $scope.ProfileQuestionnaireHeaderCode = profileQuestionnaire.Code
            $scope.ProfileQuestionnaireHeader = profileQuestionnaire;
            $scope.ProfileQuestionnaireLines = profileQuestionnaireLines;
        };

        $scope.getProfileQuestionnaireLinesAndRatingsError = function(response, status) {
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