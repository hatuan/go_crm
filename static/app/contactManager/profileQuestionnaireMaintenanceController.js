/**
 * Created by tuanha-01 on 5/13/2016.
 */
"use strict";

define(['angularAMD', 'ajaxService', 'alertsService', 'myApp.autoComplete', 'profileQuestionnairesService'], function (angularAMD, $) {
    var injectParams = ['$scope', '$rootScope', '$state', '$window', 'moment', 'alertsService', 'profileQuestionnairesService', '$stateParams', 'Constants'];

    var profileQuestionnaireMaintenanceController = function ($scope, $rootScope, $state, $window, moment, alertsService, profileQuestionnairesService, $stateParams, Constants) {

        $scope.initializeController = function () {
            $rootScope.applicationModule = "ProfileQuestionnaireMaintenance";
            $rootScope.alerts = [];

            var profileQuestionnaireID = ($stateParams.ID || "");
            
            $scope.ID = profileQuestionnaireID;
                        
            $scope.Constants = Constants;

            if (profileQuestionnaireID == "") {
                $scope.ID = "";
                $scope.Code = "";
                $scope.Description = "";
                $scope.Priority = $scope.Constants.Priorities[2].Code;
                $scope.ContactType = $scope.Constants.ContactTypes[0].Code;
                $scope.BusinessRelationTypeID = "";
                $scope.BusinessRelationTypeCode = "";

                $scope.Status = $scope.Constants.Status[1].Code;
                $scope.ClientID = "";
                $scope.OrganizationID = "";
                $scope.RecCreatedByID = $rootScope.currentUser.ID;
                $scope.RecCreatedByUser = $rootScope.currentUser.Name;
                $scope.RecCreated = new Date();
                $scope.RecModifiedByID = $rootScope.currentUser.ID;
                $scope.RecModifiedByUser = $rootScope.currentUser.Name;
                $scope.RecModified = new Date();
            } else {
                var getProfileQuestionnaire = new Object();
                getProfileQuestionnaire.ID = profileQuestionnaireID
                profileQuestionnairesService.getProfileQuestionnaire(getProfileQuestionnaire, $scope.profileQuestionnaireCompleted, $scope.profileQuestionnaireError);
            }
        };

        $scope.profileQuestionnaireCompleted = function (response, status) {
            $scope.ID = response.Data.ProfileQuestionnaire.ID;
            $scope.Code = response.Data.ProfileQuestionnaire.Code;
            $scope.Description = response.Data.ProfileQuestionnaire.Description;
            $scope.Priority = response.Data.ProfileQuestionnaire.Priority;
            $scope.ContactType = response.Data.ProfileQuestionnaire.ContactType;
            $scope.BusinessRelationTypeID = response.Data.ProfileQuestionnaire.BusinessRelationTypeID == null? "" : response.Data.ProfileQuestionnaire.BusinessRelationTypeID;
            $scope.BusinessRelationTypeCode = response.Data.ProfileQuestionnaire.BusinessRelationTypeCode;
            $scope.Status = response.Data.ProfileQuestionnaire.Status;
            $scope.Version = response.Data.ProfileQuestionnaire.Version;
            $scope.ClientID = response.Data.ProfileQuestionnaire.ClientID;
            $scope.OrganizationID = response.Data.ProfileQuestionnaire.OrganizationID;
            $scope.RecCreatedByID = response.Data.ProfileQuestionnaire.RecCreatedByID;
            $scope.RecCreatedByUser = response.Data.ProfileQuestionnaire.RecCreatedByUser;
            $scope.RecCreated = new moment.unix(response.Data.ProfileQuestionnaire.RecCreated).toDate();
            $scope.RecModifiedByID = response.Data.ProfileQuestionnaire.RecModifiedByID;
            $scope.RecModifiedByUser = response.Data.ProfileQuestionnaire.RecModifiedByUser;
            $scope.RecModified = new moment.unix(response.Data.ProfileQuestionnaire.RecModified).toDate();
        };

        $scope.profileQuestionnaireError = function (response, status) {
            alertsService.RenderErrorMessage(response.Error);
        }

        $scope.validationOptions = {
            rules: {
                Code: {
                    required: true,
                    "remote": {
                        url: "api/check-unique",
                        type: "post",
                        //dataType: 'json', //dataType is json but don't have any effect. 
                        data: {
                            UserID : function() {
                                return $rootScope.currentUser.ID
                            }, 
                            Table: "profile_questionnaire_header",
                            RecID: function() { 
                                return $scope.ID 
                            }
                        }
                    }
                },
                Description: {
                    required: true
                },
            }
        };

        $scope.update = function (form) {
            if(form.validate()) {
                var profileQuestionnaire = $scope.createProfileQuestionnaireObject();
                profileQuestionnairesService.updateProfileQuestionnaire(profileQuestionnaire, $scope.profileQuestionnaireUpdateCompleted, $scope.profileQuestionnaireUpdateError)
            }
        };

        $scope.cancel = function (form) {
           setTimeout(function() {
                $state.go('profileQuestionnaire', { profileQuestionnaireID : $scope.ID });
            }, 10);
        };

        $scope.linesMaintenance = function(form){
            if(form.validate()) {
                setTimeout(function() {
                    $state.go('profileQuestionnaireLinesMaintenance', { headerID : $scope.ID });
                }, 10);
            }
        }

        $scope.profileQuestionnaireUpdateCompleted = function (response, status) {
            $scope.ID = response.Data.ProfileQuestionnaire.ID;
            alertsService.RenderSuccessMessage(response.ReturnMessage);
            
            setTimeout(function() {
                $state.go('profileQuestionnaire', { profileQuestionnaireID : $scope.ID });
            }, 1000);
        };

        $scope.profileQuestionnaireUpdateError = function (response, status) {
            alertsService.RenderErrorMessage(response.Error);
        }

        $scope.createProfileQuestionnaireObject = function () {
            var profileQuestionnaire = new Object();
            profileQuestionnaire.ID = $scope.ID;
            profileQuestionnaire.Code = $scope.Code;
            profileQuestionnaire.Description = $scope.Description;
            profileQuestionnaire.Priority = $scope.Priority;
            profileQuestionnaire.ContactType = $scope.ContactType;
            profileQuestionnaire.BusinessRelationTypeID = $scope.BusinessRelationTypeID == "" ? null : $scope.BusinessRelationTypeID;
            profileQuestionnaire.Version = $scope.Version;

            profileQuestionnaire.Status = $scope.Status;
            profileQuestionnaire.ClientID = $scope.ClientID;
            profileQuestionnaire.OrganizationID = $scope.OrganizationID;
            profileQuestionnaire.RecCreatedByID = $scope.RecCreatedByID;
            profileQuestionnaire.RecCreatedByUser = $scope.RecCreatedByUser;
            profileQuestionnaire.RecCreated = new moment($scope.RecCreated).unix();
            profileQuestionnaire.RecModifiedByID = $rootScope.currentUser.ID;
            profileQuestionnaire.RecModifiedByUser = $rootScope.currentUser.Name;
            profileQuestionnaire.RecModified = new moment($scope.RecModified).unix();

            return profileQuestionnaire;
        }
    };

    profileQuestionnaireMaintenanceController.$inject = injectParams;
    angularAMD.controller('ProfileQuestionnaireMaintenanceController', profileQuestionnaireMaintenanceController);
});
