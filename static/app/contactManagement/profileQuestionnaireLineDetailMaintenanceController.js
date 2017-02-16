/**
 * Created by tuanha-01 on 5/13/2016.
 */
"use strict";

define(['angularAMD', 'jquery', 'ajaxService', 'alertsService', 'myApp.autoComplete', 'profileQuestionnairesService', 'app/contactManagement/ratingsMaintenanceController'], function(angularAMD, $) {
    var injectParams = ['$scope', '$rootScope', '$state', '$window', 'moment', '$uibModal', '$uibModalInstance', 'alertsService', 'profileQuestionnairesService', '$stateParams', 'Constants', 'profileQuestionnaireLine'];

    var profileQuestionnaireLineDetailMaintenanceController = function($scope, $rootScope, $state, $window, moment, $uibModal, $uibModalInstance, alertsService, profileQuestionnairesService, $stateParams, Constants, profileQuestionnaireLine) {
        $scope.Constants = Constants;
        $scope.ProfileQuestionnaireLine = profileQuestionnaireLine;

        $scope.validationOptions = {
            rules: {
                "Description[]": {
                    required: true
                }
            }
        };

        $scope.ok = function(form) {
            if (form.validate()) {
                $uibModalInstance.close($scope.ProfileQuestionnaireLine);
            }
        };

        $scope.cancel = function() {
            $uibModalInstance.dismiss('cancel');
        };

        $scope.autoContactClassificationChange = function() {

        }

        $scope.showRatingPoints = function(parentSelector) {
            var parentElem = parentSelector ?
                angular.element($document[0].querySelector('.modal-demo ' + parentSelector)) : undefined;


            var modalInstance = $uibModal.open({
                animation: true,
                ariaLabelledBy: 'modal-title',
                ariaDescribedBy: 'modal-body',
                templateUrl: 'app/contactManagement/ratingsMaintenance.html',
                controller: 'RatingsMaintenanceController',
                resolve: {
                    profileQuestionnaireLine: function() {
                        var __profileQuestionnaireLine = $.extend({}, $scope.ProfileQuestionnaireLine);
                        return __profileQuestionnaireLine;
                    },
                }
            });
            modalInstance.rendered.then(function(result) {
                $('.modal .modal-body').css('overflow-y', 'auto');
                $('.modal .modal-body').css('max-height', $(window).height() * 0.7);
                $('.modal .modal-body').css('height', $(window).height() * 0.7);
                $('.modal .modal-body').css('margin-right', 0);
            });
            modalInstance.result.then(function(editRatings) {

            }, function() {
                //dismissed 
            })['finally'](function() {
                modalInstance = undefined;
            });
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