import axios from 'axios'

export class xrpLedger {

    static getOffers(limit) {
        try {
            let ts = Math.round((new Date()).getTime() / 1000) - 50
            let url = `https://data.ripple.com/v2/transactions/?limit=${limit}&type=OfferCreate&start=${ts}`
            return axios.get(url)
        } catch (err) {
            console.log(err)
        }
    }

    static getTopCurrencies() {
        try {
            let date = new Date()
            let formatted_date = date.getFullYear() + "-" + (date.getMonth() + 1) + "-" + (date.getDate() - 1)
            let url = `https://data.ripple.com/v2/network/top_currencies/${formatted_date}?limit=10`
            return axios.get(url)
        } catch (err) {
            console.log(err)
        }
    }

    static getTopology() {
        try {
            let url = "https://data.ripple.com/v2/network/topology?verbose=true"
            return axios.get(url)
        } catch (err) {
            console.log(err)
        }
    }
}
