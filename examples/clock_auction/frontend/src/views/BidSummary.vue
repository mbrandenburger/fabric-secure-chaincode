<template>
    <div>
        <AuctionHeader></AuctionHeader>
        <BidderMenu></BidderMenu>
        <v-card>
            <v-card-title>My Bid Summary</v-card-title>
            <v-card-text class="pa-0">
                <v-data-table
                        :headers="tableHeader"
                        :items="bids"
                        hide-default-footer
                >
                    <template v-slot:item.openingPrice="props">
                        $ {{ props.item.openingPrice }}
                    </template>
                </v-data-table>
            </v-card-text>
        </v-card>
    </div>
</template>

<script>
    import { mapState } from 'vuex'

    import AuctionHeader from "../components/AppHeader";
    import BidderMenu from '../components/AppMenu'

    export default {

        name: "BidSummary",
        components: {
            AuctionHeader,
            BidderMenu,
        },
        data () {
            return {
                tableHeader: [
                    {text: 'id', value: 'id',},
                    {text: 'Round', value: 'round',},
                    {text: 'Territory', value: 'round',},
                    {text: 'Price', value: 'price',},
                    {text: 'Quantity', value: 'quantity',},
                    {text: 'Status', value: 'status',},
                ],
            }
        },
        computed: mapState({
            bids: state => state.bid.submittedBids
        }),
    }
</script>