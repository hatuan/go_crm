/**
 * Created by tuanha-01 on 5/13/2016.
 */
"use strict";

define(['angularAMD', 'jquery', 'ajaxService', 'alertsService', 'myApp.autoComplete', 'profileQuestionnairesService', 'app/contactManagement/ratingsMaintenanceController'], function(angularAMD, $) {
    var injectParams = ['$scope', '$rootScope', '$state', '$window', 'moment', '$uibModal', '$uibModalInstance', 'alertsService', 'profileQuestionnairesService', '$stateParams', '$confirm', 'Constants', 'profileQuestionnaireLine'];

    var profileQuestionnaireLineDetailMaintenanceController = function($scope, $rootScope, $state, $window, moment, $uibModal, $uibModalInstance, alertsService, profileQuestionnairesService, $stateParams, $confirm, Constants, profileQuestionnaireLine) {
        $scope.Constants = Constants;
        $scope.ProfileQuestionnaireLine = profileQuestionnaireLine;

        $scope.validationOptions = {
            rules: {
                "Description[]": {
                    required: true
                }
            }
        };

        $scope.ok = function(form, editRatings) {
            if (form.validate()) {
                var _result = new Object();
                _result.EditProfileQuestionnaireLine = $scope.ProfileQuestionnaireLine;
                _result.EditRatings = editRatings;

                $uibModalInstance.close(_result);
            }
        };

        $scope.cancel = function() {
            $uibModalInstance.dismiss('cancel');
        };

        $scope.autoContactClassificationChange = function() {
            if ($scope.ProfileQuestionnaireLine.AutoContactClassification == Constants.BooleanTypes[0].Code) { //AutoContactClassification == false
                $scope.ProfileQuestionnaireLine.CustomerClassField = Constants.ProfileQuestionaireLineCustomerClassFieldTypes[0].Code;
                $scope.ProfileQuestionnaireLine.VendorClassField = Constants.ProfileQuestionaireLineVendorClassFieldTypes[0].Code;
                $scope.ProfileQuestionnaireLine.ContactClassField = Constants.ProfileQuestionaireLineContactClassFieldTypes[0].Code;
                $scope.ProfileQuestionnaireLine.ClassificationMethod = Constants.ProfileQuestionaireLineClassificationMethodTypes[0].Code;
                $scope.ProfileQuestionnaireLine.SortingMethod = Constants.ProfileQuestionaireLineSortingMethodTypes[0].Code;
            }
        }

        $scope.disableSortingMethod = function() {
            return $scope.ProfileQuestionnaireLine.AutoContactClassification == Constants.BooleanTypes[0].Code;
        }

        $scope.disableClassificationMethod = function() {
            return $scope.ProfileQuestionnaireLine.AutoContactClassification == Constants.BooleanTypes[0].Code;
        }

        $scope.disableEndingDateFormula = function() {
            return $scope.ProfileQuestionnaireLine.AutoContactClassification == Constants.BooleanTypes[0].Code;
        }

        $scope.disableStartingDateFormula = function() {
            return $scope.ProfileQuestionnaireLine.AutoContactClassification == Constants.BooleanTypes[0].Code;
        }

        $scope.disableContactClassField = function() {
            if ($scope.ProfileQuestionnaireLine.AutoContactClassification == Constants.BooleanTypes[0].Code)
                return true;

            if ($scope.ProfileQuestionnaireLine.CustomerClassField != Constants.ProfileQuestionaireLineCustomerClassFieldTypes[0].Code)
                return true;

            if ($scope.ProfileQuestionnaireLine.VendorClassField != Constants.ProfileQuestionaireLineVendorClassFieldTypes[0].Code)
                return true;

            return false;
        }

        $scope.disableVendorClassField = function() {
            if ($scope.ProfileQuestionnaireLine.AutoContactClassification == Constants.BooleanTypes[0].Code)
                return true;

            if ($scope.ProfileQuestionnaireLine.CustomerClassField != Constants.ProfileQuestionaireLineCustomerClassFieldTypes[0].Code)
                return true;

            if ($scope.ProfileQuestionnaireLine.ContactClassField != Constants.ProfileQuestionaireLineContactClassFieldTypes[0].Code)
                return true;

            return false;
        }

        $scope.disableCustomerClassField = function() {
            if ($scope.ProfileQuestionnaireLine.AutoContactClassification == Constants.BooleanTypes[0].Code)
                return true;

            if ($scope.ProfileQuestionnaireLine.VendorClassField != Constants.ProfileQuestionaireLineVendorClassFieldTypes[0].Code)
                return true;

            if ($scope.ProfileQuestionnaireLine.ContactClassField != Constants.ProfileQuestionaireLineContactClassFieldTypes[0].Code)
                return true;

            return false;

        }
    };


    profileQuestionnaireLineDetailMaintenanceController.$inject = injectParams;
    angularAMD.controller('ProfileQuestionnaireLineDetailMaintenanceController', profileQuestionnaireLineDetailMaintenanceController);
});