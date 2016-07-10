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
            $scope.CurrentPageNumber = 1;
            $scope.SortExpression = "Code";
            $scope.SortDirection = "ASC";
            $scope.BusinessRelationTypes = [];

            $scope.getBusinessRelationTypes();
        };


        $scope.getBusinessRelationTypes = function () {
            var businessRelationTypeInquiry = $scope.createBusinessRelationTypeObject();
            businessRelationTypesService.getBusinessRelationTypes(businessRelationTypeInquiry, $scope.businessRelationTypesInquiryCompleted, $scope.businessRelationTypesInquiryError);
        };

        $scope.businessRelationTypesInquiryCompleted = function (response, status) {
            alertsService.RenderSuccessMessage(response.ReturnMessage);

            $scope.BusinessRelationTypes = response.Data.BusinessRelationTypes;
            $scope.TotalProducts = response.TotalRows;
            $scope.TotalPages = response.TotalPages;
        };

        $scope.businessRelationTypesInquiryError = function (response, status) {
            alertsService.RenderErrorMessage(response.Error);
        }


        $scope.createBusinessRelationTypeObject = function () {

            var businessRelationTypeInquiry = new Object();

            businessRelationTypeInquiry.Search = $scope.Search;
            businessRelationTypeInquiry.CurrentPageNumber = $scope.CurrentPageNumber;
            businessRelationTypeInquiry.SortExpression = $scope.SortExpression;
            businessRelationTypeInquiry.SortDirection = $scope.SortDirection;
            businessRelationTypeInquiry.PageSize = $scope.PageSize;

            return businessRelationTypeInquiry;
        }
    };

    businessRelationTypesController.$inject = injectParams;
    app.register.controller('BusinessRelationTypesController', businessRelationTypesController);
});
