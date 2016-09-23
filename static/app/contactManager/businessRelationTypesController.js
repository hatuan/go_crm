/**
 * Created by tuanha-01 on 5/13/2016.
 */
"use strict";

define(['application-configuration', 'ajaxService', 'alertsService', 'businessRelationTypesService'], function (app, $) {
    var injectParams = ['$scope', '$rootScope', '$state', '$window', 'moment', 'alertsService', 'businessRelationTypesService'];

    var businessRelationTypesController = function ($scope, $rootScope, $state, $window, moment, alertsService, businessRelationTypesService) {

        $scope.initializeController = function () {
            $rootScope.applicationModule = "BusinessRelationTypes";
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

            $scope.BusinessRelationTypes = [];
            $scope.FilteredItems = [];
            $scope.getBusinessRelationTypes();

            $scope.SearchConditions = [];
            $scope.AddSearchCondition();
        };

        $scope.refresh = function () {
            $scope.getBusinessRelationTypes();
        }

        $scope.search = function () {
            $scope.isSearched = !$scope.isSearched;
        }
        $scope.selectAll = function () {
            $scope.Selection=[];
            for(var i = 0; i < $scope.FilteredItems.length; i++){
                $scope.Selection.push($scope.FilteredItems[i]["ID"]);
            }
        }

        $scope.delete = function () {
            if($scope.Selection.length <= 0)
                return;
            var deleteBusinessRelationTypes = $scope.createDeleteBusinessRelationTypeObject()
            businessRelationTypesService.deleteBusinessRelationType(deleteBusinessRelationTypes, 
                function (response, status) {
                    $scope.getBusinessRelationTypes();
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

        $scope.getBusinessRelationTypes = function () {
            var businessRelationTypeInquiry = $scope.createBusinessRelationTypeObject();
            businessRelationTypesService.getBusinessRelationTypes(businessRelationTypeInquiry, $scope.businessRelationTypesInquiryCompleted, $scope.businessRelationTypesInquiryError);
        };

        $scope.businessRelationTypesInquiryCompleted = function (response, status) {
            alertsService.RenderSuccessMessage(response.ReturnMessage);
            $scope.BusinessRelationTypes = response.Data.BusinessRelationTypes;
            $scope.TotalRows = response.TotalRows;
            $scope.Selection = [];
            $scope.FilteredItems = [];
        };

        $scope.businessRelationTypesInquiryError = function (response, status) {
            alertsService.RenderErrorMessage(response.Error);
        }

        $scope.createBusinessRelationTypeObject = function () {
            var businessRelationTypeInquiry = new Object();

            businessRelationTypeInquiry.Search = $scope.Search;
            businessRelationTypeInquiry.SortExpression = $scope.SortExpression;
            businessRelationTypeInquiry.SortDirection = $scope.SortDirection;
            businessRelationTypeInquiry.FetchSize = $scope.FetchSize;

            return businessRelationTypeInquiry;
        }

        $scope.createDeleteBusinessRelationTypeObject = function() {
            var deleteBusinessRelationTypes = new Object();
            deleteBusinessRelationTypes.ID = $scope.Selection.join(",");
            return deleteBusinessRelationTypes;
        }

        $scope.AddSearchCondition = function(){
            var searchCondition = new Object();
            searchCondition.ID = "";
            searchCondition.Type = "FREE"; //Value in : CODE, FREE, DATE
            searchCondition.SearchValue = "BusinessRelationType";
            searchCondition.Value = "";

            $scope.SearchConditions.push(searchCondition);
        }
        
        $scope.RemoveSearchCondition = function(_index){
            $scope.SearchConditions.splice(_index, 1);
        }
    };

    businessRelationTypesController.$inject = injectParams;
    app.register.controller('BusinessRelationTypesController', businessRelationTypesController);
});
