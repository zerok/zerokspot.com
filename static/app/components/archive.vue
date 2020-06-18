<template>
    <div>
        <form v-on:submit="search">
            <p>You can select posts per year or matching a search-term here:</p>
            <select v-model="selectedYear" v-on:change="selectYear">
                <option value="">Select a year</option>
                <option v-for="year in years" v-bind:value="year[0]">{{ year[0] }} ({{ year[1] }})</option>
            </select>
            <input placeholder="Search by text..." type="search" v-model="searchTerm" />
        </form>
        <ul class="page__listing page__listing--mini" v-if="posts">
            <li v-for="post in posts" :key="post.url">
                <a v-bind:href="post.url">{{ post.title }}</a>
                <time>{{ formatDate(new Date(post.date * 1000)) }}</time>
            </li>
        </ul>
    </div>
</template>

<script language="babel">
import formatDate from 'date-fns/format';

export default {
    data() {
        return {
            years: window.years,
            selectedYear: "",
            posts: null,
            searchTerm: null
        };
    },
    created() {
        const hash = window.location.hash;
        if (hash.startsWith("#search/")) {
            const term = hash.split("#search/")[1];
            this.searchTerm = term;
            this.search();
        } else if (hash.startsWith("#year/")) {
            const year = hash.split("#year/")[1];
            this.selectedYear = year;
            this.selectYear();
        }
    },
    methods: {
        formatDate(dt) {
            return formatDate(dt, 'MMM d, yyyy');
        },
        search: function(evt) {
            if (evt) {
                evt.preventDefault();
            }
            if (!this.searchTerm) {
                return;
            }
            this.selectedYear = "";
            window.history.pushState(null, "", "#search/" + encodeURIComponent(this.searchTerm));
            fetch(window.searchBaseURL + "/search/" + encodeURIComponent(this.searchTerm)).then(resp => {
                return resp.json();
            }).then(data => {
                this.posts = data.hits;
            });
        },
        selectYear: function() {
            if (!this.selectedYear) {
                return;
            }
            this.searchTerm = null;
            window.history.pushState(null, "", "#year/" + encodeURIComponent(this.selectedYear));
            fetch(window.searchBaseURL + "/year/" + this.selectedYear).then(resp => {
                return resp.json();
            }).then(data => {
                this.posts = data.hits;
            });
        }
    }
};
</script>
