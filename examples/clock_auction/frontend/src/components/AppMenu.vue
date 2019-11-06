<template>
    <v-navigation-drawer
            :clipped=true
            :dark=true
            permanent
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
                    v-for="item in menuItems"
                    :key="item.title"
                    link
                    :to="item.link"
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
                menuItems: [
                    { title: 'Place Bids', link: '/place_bid', icon: 'fa-hammer' },
                    { title: 'Bid Summary', link: 'bid_summary', icon: 'fa-list-alt' },
                    { title: 'My Results', link: '/my_results', icon: 'fa-chart-line' },
                    { title: 'Auction Info', link: '/auction_info', icon: 'fa-info-circle' },
                    { title: 'Auction History', link: '', icon: 'fa-calendar' },
                    { title: 'Settings', link: '', icon: 'fa-cog' },
                    { title: 'Help', link: '', icon: 'fa-question-circle' },
                ],
                right: null,
            }
        },
    }
</script>