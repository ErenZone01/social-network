// Fonction CustomFetch qui prend l'URL, la méthode et les données en tant que paramètres

import sharedData from './data.js';

export async function CustomFetch(url, method, data) {
    try {
        const response = await fetch(url, {
            method: method,
            headers: {
                'Content-Type': 'application/json', // Vous pouvez ajuster les en-têtes selon vos besoins
            },
            credentials: "include",
            // N'ajoute pas le corps si la méthode est POST
            body: (method == 'POST') ? JSON.stringify(data) : undefined,
        });

        if (!response.ok) {
            throw new Error('Erreur lors de la requête fetch');
        }

        const responseData = await response.json(); // Convertit la réponse en format JSON
        return responseData; // Retourne les données récupérées
    } catch (error) {
        console.error('Erreur lors de la requête fetch :', error);
        console.error(" Error #%d", 500)
        throw error; // Lance l'erreur pour être gérée par l'appelant
    }
};

export default {
    methods: {
        async FetchCustomRef() {
            try {
                var user = parseInt(this.$route.params.userID);
                var fetch = await CustomFetch("http://localhost:8080/Profil", "POST", { ID: user });
                if (fetch.Types == "Success") {
                    sharedData.MyuserProfile = fetch.Data.Myaccount;
                    sharedData.MyProfileFollowings = fetch.Data.Allfollowing;
                    sharedData.MyProfileFollowers = fetch.Data.Allfollowers;
                    sharedData.AllId = [];
                    for (let i = 0; i < sharedData.MyProfileFollowings.length; i++) {
                        sharedData.AllId.push(sharedData.MyProfileFollowings[i].ID);
                    }
                    console.log(fetch);
                } else {
                    console.log("le message : ", fetch.Msg);
                    this.$router.push("/Login");
                }

            } catch (error) {
                console.error("Erreur lors de la récupération des données :", error);
            }
        },
        async GetData() {
            let response = await CustomFetch("http://localhost:8080/", "GET", {})
            if (response.Types == "Success") {
                console.log("success : ", response);
                sharedData.Myaccount = response.Data.Myaccount;
                sharedData.Id = sharedData.Myaccount.Id;
                sharedData.Allnotif = response.Data.Allnotif;
                sharedData.MyFollowers = response.Data.Allfollowers;
                sharedData.MyFollowings = response.Data.Allfollowing;
                sharedData.AllUsers = response.Data.Alluser;
            } else {
                console.log("error GetData : ", response.Msg);
                this.$router.push("/Login");
            }
        },
        async Decon(event) {
            event.preventDefault();
            let data = await CustomFetch("http://localhost:8080/Decon", "POST", null)
            if (data.Types == "Success") {
                console.log("vous etes deconnecté avec succée");
                this.$router.push("/Login");
            } else {
                console.log("error of logout : ", data.Msg);
            }
        },
        async Follow(user, event) {
            event.preventDefault();
            var fetch = await CustomFetch("http://localhost:8080/Follow", "POST", user);
            if (fetch.Types == "Success") {
                console.log("Follow is succes");
                sharedData.AllUsers = fetch.Data.Alluser
                sharedData.Allnotif = fetch.Data.Allnotif;
                sharedData.MyFollowings = fetch.Data.Allfollowing;
                if (this.$route.name !== 'Home') {
                    this.FetchCustomRef();
                }
                sharedData.AllId = []
                for (let i = 0; i < sharedData.MyProfileFollowings.length; i++) {
                    console.log(sharedData.MyProfileFollowings[i].ID);
                    sharedData.AllId.push(sharedData.MyProfileFollowings[i].ID);
                }
            } else {

                if (fetch.Msg.toLowerCase().includes("you have already follow")) { console.log("Error of Follow : ", fetch.Msg); } else {
                    console.log("Error of Follow : ", fetch.Msg);
                    this.$router.push("/Login");
                }
            }

            // Utilisation de this.allUsers pour faire référence à la propriété data
        },
        async Invitation(Statement, ID_Notif, event) {
            event.preventDefault();
            let response = await CustomFetch("http://localhost:8080/Invitation", "POST", { States: Statement, ID_Notif: ID_Notif });
            if (response.Types == "Success") {
                console.log("Your requete is approved");
                sharedData.Allnotif = response.Data.Allnotif;
                sharedData.MyFollowers = response.Data.Allfollowers;
            } else {
                console.log("Error Invitation : ", response.Msg);
                this.$router.push("/Login");
            }

        },
        async UnFollow(user, event) {
            event.preventDefault();
            let response = await CustomFetch("http://localhost:8080/UnFollow", "POST", user);
            if (response.Types == "Success") {
                console.log("Unfollow is success")
                sharedData.MyFollowings = response.Data;
                sharedData.AllId = []
                if (this.$route.name !== 'Home') {
                    this.FetchCustomRef();
                }
                for (let i = 0; i < sharedData.MyProfileFollowings.length; i++) {
                    console.log(sharedData.MyProfileFollowings[i].ID);
                    sharedData.AllId.push(sharedData.MyProfileFollowings[i].ID);
                }
            } else {
                console.log("Error of UnFollow : ", response.Msg);
                this.$router.push("/Login");
            }

        },
        async ChangePrivacy(e) {
            e.preventDefault();
            var fetch = await CustomFetch("http://localhost:8080/Privacy", "POST", sharedData.Myaccount);
            if (fetch.Types == "Success") {
                console.log("Change is succes");
                sharedData.Myaccount = fetch.Data.Myaccount
                sharedData.AllUsers = fetch.Data.Alluser
                sharedData.MyFollowings = fetch.Data.MyFollowings
            } else {
                console.log("Error of Change : ", fetch.Msg);
                this.$router.push("/Login");
            }
        },
        IsMyAccountFollowed() {
            console.log("je suis dedans");
            if (sharedData.MyProfileFollowers.length > 0 && sharedData.Myaccount) {
                //Vérifier si l 'ID de Myaccount est dans MyFollowers
                return sharedData.MyProfileFollowers.some(
                    follower => follower.ID === sharedData.Myaccount.ID
                );
            }
            return false;
        },
        isMyAccountFollowing() {
            if (sharedData.MyProfileFollowings && sharedData.Myaccount) {
                //Vérifier si l 'ID de Myaccount est dans MyFollowers
                return sharedData.MyProfileFollowings.some(
                    following => following.ID === sharedData.Myaccount.ID
                );

            }
            return false;
        },
        async CheckCookie() {
            var fetch = await CustomFetch("http://localhost:8080/CheckSession", "GET", null);
            if (fetch.Types == "Success") {
                console.log("status : ", fetch.Msg);
                this.$router.push("/Home");
            } else {
                console.log("status : ", fetch.Msg);
                console.log("page");
            }
        }
    },
};