<template>
    <v-navigation-drawer
            permanent
            clipped="true"
            :dark="true"
            app
    >
        <v-list-item>
            <v-list-item-content>
                <v-list-item-title class="title">
                    A-Telco
                </v-list-item-title>
                <v-list-item-subtitle>
                    A-Telco
                </v-list-item-subtitle>
            </v-list-item-content>
        </v-list-item>

        <v-divider></v-divider>

        <v-list
                dense
                nav
        >
            <v-list-item
                    v-for="item in items"
                    :key="item.title"
                    link
            >
                <v-list-item-icon>
                    <v-icon>{{ item.icon }}</v-icon>
                </v-list-item-icon>

                <v-list-item-content>
                    <v-list-item-title>{{ item.title }}</v-list-item-title>
                </v-list-item-content>
            </v-list-item>
        </v-list>

        <template v-slot:append>
            <div class="pa-2">
                <router-link to="/login">
                    <v-btn block>Logout</v-btn>
                </router-link>
            </div>
        </template>
    </v-navigation-drawer>
</template>


<script>
    export default {
        name: "BidderMenu",
        data () {
            return {
                items: [
                    { title: 'Place Bids', icon: 'fa-hammer' },
                    { title: 'Bid Summary', icon: 'fa-list-alt' },
                    { title: 'My Results', icon: 'fa-chart-line' },
                    { title: 'Auction Info', icon: 'fa-info-circle' },
                    { title: 'Auction Schedule', icon: 'fa-calendar' },
                    { title: 'Settings', icon: 'fa-cog' },
                    { title: 'Help', icon: 'fa-question-circle' },
                ],
                right: null,
            }
        },
    }
</script>