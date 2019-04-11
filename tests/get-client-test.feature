Feature: Create client integration test

  Background:
    * url 'http://localhost:8080/v1'
    * configure headers = { 'Content-Type': 'application/json' }

  Scenario Outline: create client positive
    Given request <request>
    When method post
    Then status 201
    And match response == { id:'#number' }
    And def id = response.id

    Given path id
    When method get
    Then status 200
    And match response contains <request>


    Examples:
      | request                                                                                                                                                                                                                                                       |
      | {"firstName":"TestUser","lastName":"Testerov","birthDate":"2019-01-02","Passport":{"series":1234,"number":847882},"Jobs":[{"name":"test1","wage":72364},{"name":"test4","wage":67628}],"sex":"male","maritalStatus":"married"}                                |
      | {"firstName":"TestUser","lastName":"Testerov","birthDate":"2005-05-30","Passport":{"series":4356,"number":351367},"Jobs":[{"name":"test1","wage":187623},{"name":"test4","wage":32641}],"sex":"male","maritalStatus":"single","children":0}                   |
      | {"firstName":"TestUser","lastName":"Testerov","birthDate":"2012-04-03","Passport":{"series":9847,"number":866234},"Jobs":[{"name":"test1","wage":36434},{"name":"test4","wage":4637}],"sex":"female","maritalStatus":"married","children":1}                  |
      | {"firstName":"TestUser","lastName":"Testerov","birthDate":"2007-07-07","Passport":{"series":2345,"number":215344},"Jobs":[],"sex":"male","maritalStatus":"married","children":2}                                                                              |
      | {"firstName":"TestUser","lastName":"Testerov","birthDate":"2000-01-01","Passport":{"series":2342,"number":812636},"Jobs":[{"name":"market","wage":23667}],"sex":"male","maritalStatus":"married","children":3}                                                |
      | {"firstName":"TestUser","lastName":"Testerov","birthDate":"1999-12-31","Passport":{"series":4361,"number":465283},"Jobs":[{"name":"test1","wage":19437},{"name":"test4","wage":23856},{"name":"test4","wage":23565}],"sex":"female","maritalStatus":"single"} |
      | {"firstName":"TestUser","lastName":"Testerov","birthDate":"1899-05-05","Passport":{"series":9730,"number":198243},"Jobs":[{"name":"test1","wage":536446},{"name":"test4","wage":53644346}],"sex":"male","maritalStatus":"married","children":2}               |
      | {"firstName":"TestUser","lastName":"Testerov","birthDate":"2000-09-04","Passport":{"series":2342,"number":210983},"Jobs":[{"name":"test1","wage":45638},{"name":"test4","wage":23654}],"sex":"male","maritalStatus":"married","children":3}                   |
      | {"firstName":"TestUser","lastName":"Testerov","birthDate":"1015-03-27","Passport":{"series":8763,"number":328654},"Jobs":[{"name":"test1","wage":36473},{"name":"test4","wage":43645}],"sex":"female","maritalStatus":"married","children":1}                 |
      | {"firstName":"TestUser","lastName":"Testerov","birthDate":"2001-08-23","Passport":{"series":6125,"number":328742},"Jobs":[{"name":"test1","wage":438765},{"name":"test4","wage":164425}],"sex":"male","maritalStatus":"married"}                              |
      | {"firstName":"TestUser","lastName":"Testerov","birthDate":"2010-01-12","Passport":{"series":2374,"number":934732},"Jobs":[{"name":"test1","wage":12873},{"name":"test4","wage":62452}],"sex":"female","maritalStatus":"married","children":4}                 |


  Scenario Outline: create credit negative
    Given path <ID>
    When method get
    And match response contains any { message:'#string',description:'#list'}

    Examples:
      | ID          |
      | 43664543    |
      | "dfbfbgbgg" |
      | -3453       |
      | "null"      |

