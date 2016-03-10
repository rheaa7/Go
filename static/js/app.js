"use strict";

angular.module('SearchWords', [])
	.controller('SearchController', function($scope, $http) {
		$scope.searchString = function(search, max) {

			var data = {
				search : search,
				max : max
			}
			console.log("Search:" + search)
			
			if (max >= 10 && max <= 50) {
				$http.get('/api/v1/suggestions/?search=' + search + '&max=' + max)
					.then(function(results) {
						console.log(results)
						console.log(results.data)
						$scope.results = results.data.word;
						
					}).catch(function(err) {
						if (err == 400) {
							$('#400').css('visibility', 'visible');
						}
					})
				} 
		}
	})