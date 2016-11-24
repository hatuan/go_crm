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

            $scope.ProfileQuestionnaireLines = response.Data.ProfileQuestionnaireLines;
            $scope.TotalRows = response.TotalRows;
        };

        $scope.getProfileQuestionnaireLinesError = function (response, status) {
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
                
            }
        };

        $scope.cancel = function (form) {
           setTimeout(function() {
                $state.go('profileQuestionnaireMaintenance', { ID : $scope.profileQuestionnaireHeaderID });
            }, 10);
        };
        
        $scope.addLines = function(){
            var profileQuestionnaireLine = $scope.createProfileQuestionnaireLineObject();
            $scope.ProfileQuestionnaireLines.push(profileQuestionnaireLine); 
        }

        $scope.lineDetail = function(profileQuestionnaireLine){
             
        }

        $scope.createProfileQuestionnaireLineObject = function () {
            var profileQuestionnaireLine = new Object();
            profileQuestionnaireLine.ID = "";
            profileQuestionnaireLine.Type = $scope.Constants.ProfileQuestionaireLineTypes[0].Code;
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
            profileQuestionnaireLine.FromValue = "";
            profileQuestionnaireLine.ToValue = "";

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
