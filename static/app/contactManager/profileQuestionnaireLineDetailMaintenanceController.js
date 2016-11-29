/**
 * Created by tuanha-01 on 5/13/2016.
 */
"use strict";

define(['angularAMD', 'ajaxService', 'alertsService', 'myApp.autoComplete', 'profileQuestionnairesService'], function (angularAMD, $) {
    var injectParams = ['$scope', '$rootScope', '$state', '$window', 'moment', '$uibModalInstance', 'alertsService', 'profileQuestionnairesService', '$stateParams', 'Constants', 'profileQuestionnaireLine'];

    var profileQuestionnaireLineDetailMaintenanceController = function ($scope, $rootScope, $state, $window, moment, $uibModalInstance, alertsService, profileQuestionnairesService, $stateParams, Constants, profileQuestionnaireLine) {

        $scope.initializeController = function () {
            $scope.Constants = Constants;
            $scope.ProfileQuestionnaireLine = profileQuestionnaireLine;

        };

        $scope.validationOptions = {
            rules: {
                "Description[]": {
                    required: true
                }
            }
        };

        $scope.ok = function () {
            $uibModalInstance.close($scope.ProfileQuestionnaireLine);
        };

        $scope.cancel = function () {
            $uibModalInstance.dismiss('cancel');
        };

        $scope.autoContactClassificationChange = function() {

        }
    };


    profileQuestionnaireLineDetailMaintenanceController.$inject = injectParams;
    angularAMD.controller('profileQuestionnaireLineDetailMaintenanceController', profileQuestionnaireLineDetailMaintenanceController);
});
