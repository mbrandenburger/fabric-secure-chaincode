<template>
    <div>
        <AuctionHeader></AuctionHeader>
        <BidderMenu></BidderMenu>
        <ClockBidding v-if="auction.currentPhase === 1"></ClockBidding>
        <AssignmentBidding v-else></AssignmentBidding>
    </div>
</template>

<script>

    import { mapState } from 'vuex'
    import AuctionHeader from "../components/AppHeader"
    import BidderMenu from '../components/AppMenu'
    import ClockBidding from '../components/Bidding/ClockBidding'
    import AssignmentBidding from '../components/Bidding/AssignmentBidding'

    export default {

        name: "PlaceBid",
        components: {
            AuctionHeader,
            BidderMenu,
            ClockBidding,
            AssignmentBidding,
        },
        computed: mapState({
            auction: state => state.auction.auction
        })
    }
</script>