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
            
            $scope.multiple = true;
            $scope.auto = true;
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
                var profileQuestionnaire = $scope.createProfileQuestionnaireObject();
                profileQuestionnairesService.updateProfileQuestionnaire(profileQuestionnaire, $scope.profileQuestionnaireUpdateCompleted, $scope.profileQuestionnaireUpdateError)
            }
        };

        $scope.cancel = function (form) {
           setTimeout(function() {
                $state.go('profileQuestionnaireMaintenance', { ID : $scope.profileQuestionnaireHeaderID });
            }, 10);
        };
 
    };

    profileQuestionnaireLinesMaintenanceController.$inject = injectParams;
    angularAMD.controller('ProfileQuestionnaireLinesMaintenanceController', profileQuestionnaireLinesMaintenanceController);
});
