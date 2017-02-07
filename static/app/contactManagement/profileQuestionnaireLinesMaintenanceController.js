/**
 * Created by tuanha-01 on 5/13/2016.
 */
"use strict";

define(['angularAMD', 'jquery', 'ajaxService', 'alertsService', 'myApp.autoComplete', 'profileQuestionnairesService', 'app/contactManagement/profileQuestionnaireLineDetailMaintenanceController'], function (angularAMD, $) {
    var injectParams = ['$scope', '$rootScope', '$state', '$window', 'moment', '$uibModal', 'alertsService', 'profileQuestionnairesService', '$stateParams', 'Constants'];

    var profileQuestionnaireLinesMaintenanceController = function ($scope, $rootScope, $state, $window, moment, $uibModal, alertsService, profileQuestionnairesService, $stateParams, Constants) {

        $scope.initializeController = function () {
            $rootScope.applicationModule = "ProfileQuestionnaireLinesMaintenance";
            $rootScope.alerts = [];

            $scope.profileQuestionnaireHeaderID = ($stateParams.headerID || "");

            $scope.Constants = Constants;

            $scope.ProfileQuestionnaireLines = [];
            $scope.ProfileQuestionnaireLineDeletes = [];

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
            $scope.ProfileQuestionnaireLineDeletes = [];
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
                "FromValue[]": {
                    number: true
                },
                "ToValue[]": {
                    number: true
                }
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

                    if (angular.isUndefinedOrNull(profileQuestionnaireLines[i].FromValue) || profileQuestionnaireLines[i].FromValue == "")
                        profileQuestionnaireLines[i].FromValue = "0";
                    if (angular.isUndefinedOrNull(profileQuestionnaireLines[i].ToValue) || profileQuestionnaireLines[i].ToValue == "")
                        profileQuestionnaireLines[i].ToValue = "0";
                }

                profileQuestionnairesService.updateProfileQuestionnaireLines(profileQuestionnaire, profileQuestionnaireLines, $scope.profileQuestionnaireLinesUpdateCompleted, $scope.profileQuestionnaireLinesUpdateError)
            }
        };

        $scope.profileQuestionnaireLinesUpdateCompleted = function (response, status) {
            alertsService.RenderSuccessMessage(response.ReturnMessage);

            setTimeout(function () {
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

        $scope.detailLine = function (_profileQuestionnaireLine, parentSelector) {
            var parentElem = parentSelector ?
                angular.element($document[0].querySelector('.modal-demo ' + parentSelector)) : undefined;

                
            var modalInstance = $uibModal.open({
                animation: true,
                ariaLabelledBy: 'modal-title',
                ariaDescribedBy: 'modal-body',
                templateUrl: 'app/contactManagement/profileQuestionnaireLineDetailMaintenance.html',
                controller: 'profileQuestionnaireLineDetailMaintenanceController',
                resolve: {
                    profileQuestionnaireLine: function () {
                        var __profileQuestionnaireLine = $.extend({}, _profileQuestionnaireLine) ;
                        return __profileQuestionnaireLine;
                    }
                }
            });
            modalInstance.rendered.then(function (result) {
                $('.modal .modal-body').css('overflow-y', 'auto');
                $('.modal .modal-body').css('max-height', $(window).height()*0.7);
                $('.modal .modal-body').css('height', $(window).height()*0.7);
                $('.modal .modal-body').css('margin-right', 0);
            });
            modalInstance.result.then(function (editProfileQuestionnaireLine) {
                _profileQuestionnaireLine.Description = editProfileQuestionnaireLine.Description;
                _profileQuestionnaireLine.MultipleAnswers = editProfileQuestionnaireLine.MultipleAnswers;
                _profileQuestionnaireLine.AutoContactClassification = editProfileQuestionnaireLine.AutoContactClassification;
                _profileQuestionnaireLine.CustomerClassField = editProfileQuestionnaireLine.CustomerClassField;
                _profileQuestionnaireLine.VendorClassField = editProfileQuestionnaireLine.VendorClassField;
                _profileQuestionnaireLine.ContactClassField = editProfileQuestionnaireLine.ContactClassField;
                _profileQuestionnaireLine.StartingDateFormula = editProfileQuestionnaireLine.StartingDateFormula;
                _profileQuestionnaireLine.EndingDateFormula = editProfileQuestionnaireLine.EndingDateFormula;
                _profileQuestionnaireLine.ClassificationMethod = editProfileQuestionnaireLine.ClassificationMethod;
                _profileQuestionnaireLine.SortingMethod = editProfileQuestionnaireLine.SortingMethod;

            }, function () {
                //dismissed 
            })['finally'](function () {
                modalInstance = undefined  // <--- This fixes
            });
            
        }

        $scope.insertLine = function (_profileQuestionnaireLine, _index) {
            var _insertProfileQuestionnaireLine = $scope.createProfileQuestionnaireLineObject();
            _insertProfileQuestionnaireLine.LineNo = _index;

            for(var _i = _index; _i <= $scope.ProfileQuestionnaireLines.length; _i++) {
                $scope.ProfileQuestionnaireLines[_i - 1].LineNo = _i + 1;
            }
            $scope.ProfileQuestionnaireLines.splice(_index - 1, 0, _insertProfileQuestionnaireLine);
        }

        $scope.removeLine = function (_profileQuestionnaireLine, _index) {
            //TODO: Need check Profile Questionnaire Line is used before delete

            for(var _i = _index + 1; _i <= $scope.ProfileQuestionnaireLines.length; _i++) {
                $scope.ProfileQuestionnaireLines[_i - 1].LineNo = _i - 1;
            }

            $scope.ProfileQuestionnaireLines.splice(_index - 1, 1);
        }

        $scope.createProfileQuestionnaireLineObject = function () {
            var profileQuestionnaireLine = new Object();
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
            profileQuestionnaireLine.FromValue = "0";
            profileQuestionnaireLine.ToValue = "0";

            profileQuestionnaireLine.Status = $scope.Constants.Status[1].Code;
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
