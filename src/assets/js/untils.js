// Dans votre fichier untils.js

// Fonction CustomFetch qui prend l'URL, la méthode et les données en tant que paramètres
export async function CustomFetch(url, method, data) {
    try {
        const response = await fetch(url, {
            method: method,
            headers: {
                'Content-Type': 'application/json', // Vous pouvez ajuster les en-têtes selon vos besoins
            },
            body: JSON.stringify(data), // Convertit les données en format JSON,
            credentials: "include"
        });

        if (!response.ok) {
            throw new Error('Erreur lors de la requête fetch');
        }

        const responseData = await response.json(); // Convertit la réponse en format JSON
        return responseData; // Retourne les données récupérées
    } catch (error) {
        console.error('Erreur lors de la requête fetch :', error);
        throw error; // Lance l'erreur pour être gérée par l'appelant
    }
}

import sharedData from './data.js';

export default {
    methods: {
        async FetchCustomRef() {
            try {
                var user = parseInt(this.$route.params.userID);
                var fetch = await CustomFetch("http://localhost:8080/Profil", "POST", { ID: user });
                sharedData.Myuser = fetch.Data.Myaccount;
                sharedData.MyFollowings = fetch.Data.Allfollowing;
                sharedData.MyFollowers = fetch.Data.Allfollowers;
                console.log(fetch);
            } catch (error) {
                console.error("Erreur lors de la récupération des données :", error);
            }
        },
        GetData() {
            fetch("http://localhost:8080/", {
                    method: "GET",
                    headers: { "Content-type": "Application/Json" },
                    credentials: "include",
                })
                .then((responses) => responses.json())
                .then((response) => {
                    if (response.Types == "Success") {
                        console.log("success : ", response);
                        sharedData.Myaccount = response.Data.Myaccount;
                        console.log("l'id est : ", sharedData.Myaccount.ID);
                        sharedData.Id = sharedData.Myaccount.Id;
                        sharedData.Allnotif = response.Data.Allnotif;
                        sharedData.MyFollowers = response.Data.Allfollowers;
                        sharedData.MyFollowings = response.Data.Allfollowing;
                        sharedData.AllUsers = response.Data.Alluser;
                    } else {
                        console.log("error l'utilisateur n'est pas connecté");
                        this.$router.push("/Login");
                    }
                    // Utilisation de this.allUsers pour faire référence à la propriété data
                })
                .catch((err) => console.log(err));
        },
        Decon(event) {
            event.preventDefault();
            fetch("http://localhost:8080/Decon", {
                    method: "POST",
                    headers: { "Content-type": "Application/Json" },
                    credentials: "include",
                })
                .then((response) => response.json())
                .then((data) => {
                    if (data.Types == "Success") {
                        console.log("vous etes deconnecté avec succée");
                        this.$router.push("/Login");
                    } else {
                        console.log(data.Msg);
                    }
                })
                .catch((error) => {
                    console.log("error lors de la deconnection : ", error);
                });
        },
        Follow(user, event) {
            event.preventDefault();
            fetch("http://localhost:8080/Follow", {
                    method: "POST",
                    body: JSON.stringify(user),
                    headers: { "Content-type": "Application/Json" },
                    credentials: "include",
                })
                .then((responses) => responses.json())
                .then((response) => {
                    console.log("Follow is succes");
                    sharedData.AllUsers = response.Data.Alluser
                    sharedData.Allnotif = response.Data.Allnotif;
                    sharedData.MyFollowings = response.Data.Allfollowing;
                    // Utilisation de this.allUsers pour faire référence à la propriété data
                })
                .catch((err) => console.log(err));
        },
        Invitation(Statement, ID_Notif, event) {
            event.preventDefault();
            fetch("http://localhost:8080/Invitation", {
                    method: "POST",
                    body: JSON.stringify({ States: Statement, ID_Notif: ID_Notif }),
                    headers: { "Content-type": "Application/Json" },
                    credentials: "include",
                })
                .then((responses) => responses.json())
                .then((response) => {
                    if (response.Types == "Success") {
                        console.log("Your requete is approved");
                        sharedData.Allnotif = response.Data.Allnotif;
                        sharedData.MyFollowers = response.Data.Allfollowers;
                    } else {
                        console.log(response.Msg);
                    }
                })
                .catch((error) => {
                    console.log("error lors de l'invitation :  ", error);
                });
        },
        UnFollow(user, event) {
            event.preventDefault();
            fetch("http://localhost:8080/UnFollow", {
                    method: "POST",
                    body: JSON.stringify(user),
                    headers: { "Content-type": "Application/Json" },
                    credentials: "include",
                })
                .then((responses) => responses.json())
                .then((response) => {
                    console.log("Unfollow is success")
                    sharedData.MyFollowings = response.Data;
                    // Utilisation de this.allUsers pour faire référence à la propriété data
                })
                .catch((err) => console.log(err));
        },
    },
};