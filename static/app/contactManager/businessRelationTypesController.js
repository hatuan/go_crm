/**
 * Created by tuanha-01 on 5/13/2016.
 */
"use strict";

define(['application-configuration', 'ajaxService', 'alertsService', 'sqlParseService', 'businessRelationTypesService'], function (app, $) {
    var injectParams = ['$scope', '$rootScope', '$state', '$window', 'moment', 'alertsService', 'sqlParseService', 'businessRelationTypesService'];

    var businessRelationTypesController = function ($scope, $rootScope, $state, $window, moment, alertsService, sqlParseService, businessRelationTypesService) {

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

            $scope.SearchConditionObjects = [];
            $scope.SearchConditionObjects.push({
                'ID': "business_relation_type.code",
                'Name': "Code",
                'Type': "CODE", //CODE, FREE, DATE
                'ValueIn': "BusinessRelationType",
                'Value': ""
            },
            {
                'ID': "business_relation_type.name",
                'Name': "Name",
                'Type': "FREE", //CODE, FREE, DATE
                'ValueIn': "",
                'Value': ""
            });

            $scope.SearchConditions = [];
            $scope.AddSearchCondition();

            $scope.BusinessRelationTypes = [];
            $scope.FilteredItems = [];
            $scope.getBusinessRelationTypes();
        };

        $scope.refresh = function () {
            $scope.getBusinessRelationTypes();
        }

        $scope.showSearch = function () {
            $scope.isSearched = !$scope.isSearched;
        }

        $scope.startSearch = function() {
            var searchConditions = [];
            
            for(var _i = 0; _i < $scope.SearchConditions.length; _i ++) {
                searchConditions.push({
                    "ID": $scope.SearchConditions[_i].Object.ID,
                    "Value": $scope.SearchConditions[_i].Object.Value,
                });
            }

            sqlParseService.getSqlCondition(searchConditions, $scope.searchCompleted, $scope.searchError);
        }

        $scope.searchCompleted = function(response, status) {
            var errs = response.Data.Errs;
            var stmts = response.Data.Stmts;
             
            for(var _i = 0; _i < errs.length; _i ++) {
                $scope.SearchConditions[_i].Error = "";
                $scope.SearchConditions[_i].HasError = false;
                $scope.SearchConditions[_i].Stmt = stmts[_i];
                if (_i == 0)
                    $scope.Search = $scope.SearchConditions[_i].Stmt;
                else 
                    $scope.Search += "AND (" + $scope.SearchConditions[_i].Stmt + ")";
            }
            $scope.getBusinessRelationTypes();
        }

        $scope.searchError = function(response, status) {
            var errs = response.Data.Errs;
            var stmts = response.Data.Stmts;
            for(var _i = 0; _i < errs.length; _i ++) {
                $scope.SearchConditions[_i].Error = errs[_i];
                $scope.SearchConditions[_i].HasError = errs[_i].length > 0;
                $scope.SearchConditions[_i].Stmt = stmts[_i];
            }
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
            searchCondition.No = $scope.SearchConditions.length + 1;
            searchCondition.Err = "";
            searchCondition.HasErr = false;
            searchCondition.Stmt = "";
            searchCondition.Object = JSON.parse(JSON.stringify($scope.SearchConditionObjects[0]));

            $scope.SearchConditions.push(searchCondition);
        }
        
        $scope.RemoveSearchCondition = function(_index){
            $scope.SearchConditions.splice(_index, 1);
        }
    };

    businessRelationTypesController.$inject = injectParams;
    app.register.controller('BusinessRelationTypesController', businessRelationTypesController);
});
