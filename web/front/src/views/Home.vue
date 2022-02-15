<template>
  <div>
     <v-app app>
        <TopBar></TopBar>
        <!-- <router-view  :key="$router.fullpath"></router-view> -->

        <v-main class="grey lighten-3">
          <v-container>
            <v-row>
                <v-col cols="3"  >
                  <Nav></Nav>
                </v-col>
                <v-col cols="9">
                  <v-sheet max-width="60vw" min-height="80vh" rounded="lg">
                    <router-view  :key="$router.path"></router-view>
                    <!-- <Articlelist></Articlelist> -->
                  </v-sheet>
                </v-col>
            </v-row>
          </v-container>
        </v-main>

        <Footer></Footer>
     </v-app>
  </div>
</template>

<script>
import TopBar from '../components/TopBar.vue'
import Footer from '../components/Footer.vue'
import Nav from '../components/Nav.vue'
import Articlelist from "../components/ArticleList.vue"
  export default {
    components: {
      TopBar,
      Footer,
      Nav,
      Articlelist
    },
  }
</script>
