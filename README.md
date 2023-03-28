# InstaSafe_Coding_Assesment_without_databse
Data is only locally stored, so if the server is restarted, the data will be permanently deleted.

# Run the code by using the following commands:

--> cd server

--> go run main.go

# API

--> Import "PostmanCollection.json" file to Postman.

1-> Create_End_User : POST
    
    url: http://localhost:5056/user/signUp

    sampleInput :   {
                        "Username":"Shashank",
                        "Email":"Shashank@gmail.com"
                    }


2-> Add_Transaction : POST

    url: http://localhost:5056/transactions

    sampleInput :   {
                        "amount":"10.23",
                        "end_user_name":"Shashank",
                        "end_user_email":"Shashank@gmail.com",
                        "timestamp":"2023-03-27T16:14:22.194962Z",
                        "location":"Bangalore"
                    }


3-> Add_Loaction : POST

    url: http://localhost:5056/user/:uid/addLoaction  

         (http://localhost:5056/user/6421bfb857da685452441912/addLoaction )

    sampleInput :   {
                        "city":"Mangalore"
                    }

4-> Reset_Loaction : PUT

    url: http://localhost:5056/user/:uid/resetLoaction

        (http://localhost:5056/user/6421bfb857da685452441912/resetLoaction)

    sampleInput :   {
                        "city":"Udupi"
                    }

5-> Get_Statistics : GET

    url: http://localhost:5056/user/:uid/statistics

         (http://localhost:5056/user/6421bfb857da685452441912/statistics?city=Mangalore)


6-> Delete_All_Transactions : DELETE

    url: http://localhost:5056/transactions

        