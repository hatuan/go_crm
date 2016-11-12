/**
 * Created by tuanha-01 on 5/13/2016.
 */
"use strict";

define(['angularAMD', 'ajaxService', 'alertsService', 'myApp.Search', 'profileQuestionnairesService'], function (angularAMD) {
    var injectParams = ['$scope', '$rootScope', '$state', '$window', 'moment', 'alertsService', 'profileQuestionnairesService'];

    var profileQuestionnairesController = function ($scope, $rootScope, $state, $window, moment, alertsService, profileQuestionnairesService) {

        $scope.initializeController = function () {
            $rootScope.applicationModule = "ProfileQuestionnaires";
            $rootScope.alerts = [];

            $scope.Search = "";
            $scope.isSearched = false;
            $scope.SortExpression = "Code";
            $scope.SortDirection = "ASC";
            $scope.FetchSize = 100;
            $scope.CurrentPage = 1;
            $scope.PageSize = 9;
            $scope.TotalRows = 0;
            $scope.Selection=[];

            $scope.searchConditionObjects = [];
            $scope.searchConditionObjects.push({
                ID: "profile_questionnaire_header.code",
                Name: "Code",
                Type: "CODE", //CODE, FREE, DATE
                ValueIn: "ProfileQuestionnaire",
                Value: ""
            },
            {
                ID: "profile_questionnaire_header.description",
                Name: "Description",
                Type: "FREE", //CODE, FREE, DATE
                ValueIn: "",
                Value: ""
            });

            $scope.ProfileQuestionnaires = [];
            $scope.FilteredItems = [];
            $scope.getProfileQuestionnaires();
        };

        $scope.refresh = function () {
            $scope.getProfileQuestionnaires();
        }

        $scope.showSearch = function () {
            $scope.isSearched = !$scope.isSearched;
        }

        $scope.selectAll = function () {
            $scope.Selection=[];
            for(var i = 0; i < $scope.FilteredItems.length; i++) {
                $scope.Selection.push($scope.FilteredItems[i]["ID"]);
            }
        }

        $scope.delete = function () {
            if($scope.Selection.length <= 0)
                return;
            var deleteProfileQuestionnaires = $scope.createDeleteProfileQuestionnaireObject()
            profileQuestionnairesService.deleteProfileQuestionnaire(deleteProfileQuestionnaires, 
                function (response, status) {
                    $scope.getProfileQuestionnaires();
                }, 
                function (response, status){
                    alertsService.RenderErrorMessage(response.Error);
                });    
        }

        $scope.toggleSelection = function (_id) {
             var idx = $scope.Selection.indexOf(_id);
             if (idx > -1) {
               $scope.Selection.splice(idx, 1);
             }
             else {
               $scope.Selection.push(_id);
             }
        };

        $scope.getProfileQuestionnaires = function (searchSqlCondition) {
            if(!angular.isUndefinedOrNull(searchSqlCondition))
                $scope.Search = searchSqlCondition;
            var profileQuestionnaireInquiry = $scope.createProfileQuestionnaireObject();
            profileQuestionnairesService.getProfileQuestionnaires(profileQuestionnaireInquiry, $scope.profileQuestionnairesInquiryCompleted, $scope.profileQuestionnairesInquiryError);
        };

        $scope.profileQuestionnairesInquiryCompleted = function (response, status) {
            alertsService.RenderSuccessMessage(response.ReturnMessage);
            $scope.ProfileQuestionnaires = response.Data.ProfileQuestionnaires;
            $scope.TotalRows = response.TotalRows;
            $scope.Selection = [];
            $scope.FilteredItems = [];
        };

        $scope.profileQuestionnairesInquiryError = function (response, status) {
            alertsService.RenderErrorMessage(response.Error);
        }

        $scope.createProfileQuestionnaireObject = function () {
            var profileQuestionnaireInquiry = new Object();

            profileQuestionnaireInquiry.Search = $scope.Search;
            profileQuestionnaireInquiry.SortExpression = $scope.SortExpression;
            profileQuestionnaireInquiry.SortDirection = $scope.SortDirection;
            profileQuestionnaireInquiry.FetchSize = $scope.FetchSize;

            return profileQuestionnaireInquiry;
        }

        $scope.createDeleteProfileQuestionnaireObject = function() {
            var deleteProfileQuestionnaires = new Object();
            deleteProfileQuestionnaires.ID = $scope.Selection.join(",");
            return deleteProfileQuestionnaires;
        }
    };

    profileQuestionnairesController.$inject = injectParams;
    angularAMD.controller('ProfileQuestionnairesController', profileQuestionnairesController);
});
