"use strict";

angular.module('SearchWords', [])
	.controller('SearchController', function($scope, $http) {
		$scope.searchString = function(search, max) {
			var data = {
				search : search,
				max : max
			}
			console.log(data)
			
			if (max >= 10) {
				$http.get('/api/v1/suggestions/?search=' + search + '&max=' + max)
					.then(function(results) {
						console.log(results)
						console.log(results.data)
						$scope.results = results.data.word;
						
					})
				} 
		}
	})