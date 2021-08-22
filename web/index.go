package web

// IndexTmpl html template for the start page
const IndexTmpl = `
<!DOCTYPE html>
<html>
  <head>
    <title>Free Blocky Listener Daemon</title>
    <meta charset="utf-8">
      <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
      <meta name="description" content="Free Blocky Listener Daemon - Safety internet for adult, children, students and small organizations.">
      <meta name="author" content="Yevgeniy Goncharov, https://sys-adm.in">
      <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.3/css/bulma.min.css">
      <script src="https://lab.sys-adm.in/js/main.js"></script>
  </head>
  
    <body>
      <!--      # Header-->
        <section class="section pt-0 has-background-light">
          <nav class="navbar py-4" style="background-color: transparent;">
            <div class="navbar-brand">
              <a class="navbar-item" href="https://lab.sys-adm.in" target="_blank">
                <img class="image" src="https://lab.sys-adm.in/images/lab.sys-adm.in-logo.png" alt="" width="96px">
              </a>
              <a class="navbar-burger" role="button" aria-label="menu" aria-expanded="false">
                <span aria-hidden="true"></span>
                <span aria-hidden="true"></span>
                <span aria-hidden="true"></span>
              </a>
            </div>
            <div class="navbar-menu">
              <div class="navbar-end">
                <a class="navbar-item" href="https://github.com/0xERR0R/blocky" target="_blank">Based on Blocky</a>
                <a class="navbar-item" href="https://sys-adm.in" target="_blank">Author</a>
              </div>
            </div>
          </nav>
          <div class="container pt-5">
            <div class="columns is-multiline is-centered mb-6">
              <div class="column is-8 is-6-desktop">
                <div class="has-text-centered">
                  <span class="has-text-grey-dark">Focus on information for free</span>
                  <h2 class="mt-3 mb-5 is-size-2 is-size-3-mobile has-text-weight-bold">
                    Web without tracking, advertising, malware and etc with<br />
                    Blocky Listener Daemon<br/>from Sys-Admin
                  </h2>
                  <p class="subtitle">You can use BLD as:</p>
                  <p>DoH (in browsers):<code>https://bld.sys-adm.in:8443/dns-query</code></p>
                  <p>DoT (as private DNS in Android):<code>bld.sys-adm.in</code></p>
                  <p>DNS (any IP from):<code>nslookup bld.sys-adm.in</code></p>
                </div>
              </div>
            </div>
            <div class="columns is-multiline">
              <div class="column is-4 mb-5 is-flex">
                <div class="mr-4">
                  <svg viewbox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg" style="width: 48px;height: 48px">
                    <path d="M25.6 22.9C25.7 23 25.8 23 26 23H33C33.6 23 34 22.6 34 22C34 21.8 34 21.7 33.9 21.6L30.4 14.6C30.1 14.1 29.5 13.9 29 14.2C28.9 14.3 28.7 14.4 28.6 14.6L25.1 21.6C24.9 22 25.1 22.6 25.6 22.9ZM29.5 17.2L31.4 21H27.6L29.5 17.2ZM18.5 14C16 14 14 16 14 18.5C14 21 16 23 18.5 23C21 23 23 21 23 18.5C23 16 21 14 18.5 14ZM18.5 21C17.1 21 16 19.9 16 18.5C16 17.1 17.1 16 18.5 16C19.9 16 21 17.1 21 18.5C21 19.9 19.9 21 18.5 21ZM22.7 25.3C22.3 24.9 21.7 24.9 21.3 25.3L18.5 28.1L15.7 25.3C15.3 24.9 14.7 24.9 14.3 25.3C13.9 25.7 13.9 26.3 14.3 26.7L17.1 29.5L14.3 32.3C13.9 32.7 13.9 33.3 14.3 33.7C14.7 34.1 15.3 34.1 15.7 33.7L18.5 30.9L21.3 33.7C21.7 34.1 22.3 34.1 22.7 33.7C23.1 33.3 23.1 32.7 22.7 32.3L19.9 29.5L22.7 26.7C23.1 26.3 23.1 25.7 22.7 25.3ZM33 25H26C25.4 25 25 25.4 25 26V33C25 33.6 25.4 34 26 34H33C33.6 34 34 33.6 34 33V26C34 25.4 33.6 25 33 25ZM32 32H27V27H32V32Z" fill="#00d1b2"></path>
                    <circle cx="24" cy="24" r="23.5" stroke="#00d1b2"></circle>
                  </svg>
                </div>
                <div>
                  <h4 class="mb-2 is-size-5 has-text-weight-bold">Fast pages loading</h4>
                  <p>Without advertising and trackers</p>
                </div>
              </div>
              <div class="column is-4 mb-5 is-flex">
                <div class="mr-4">
                  <svg viewbox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg" style="width: 48px;height: 48px">
                    <path d="M25.6 22.9C25.7 23 25.8 23 26 23H33C33.6 23 34 22.6 34 22C34 21.8 34 21.7 33.9 21.6L30.4 14.6C30.1 14.1 29.5 13.9 29 14.2C28.9 14.3 28.7 14.4 28.6 14.6L25.1 21.6C24.9 22 25.1 22.6 25.6 22.9ZM29.5 17.2L31.4 21H27.6L29.5 17.2ZM18.5 14C16 14 14 16 14 18.5C14 21 16 23 18.5 23C21 23 23 21 23 18.5C23 16 21 14 18.5 14ZM18.5 21C17.1 21 16 19.9 16 18.5C16 17.1 17.1 16 18.5 16C19.9 16 21 17.1 21 18.5C21 19.9 19.9 21 18.5 21ZM22.7 25.3C22.3 24.9 21.7 24.9 21.3 25.3L18.5 28.1L15.7 25.3C15.3 24.9 14.7 24.9 14.3 25.3C13.9 25.7 13.9 26.3 14.3 26.7L17.1 29.5L14.3 32.3C13.9 32.7 13.9 33.3 14.3 33.7C14.7 34.1 15.3 34.1 15.7 33.7L18.5 30.9L21.3 33.7C21.7 34.1 22.3 34.1 22.7 33.7C23.1 33.3 23.1 32.7 22.7 32.3L19.9 29.5L22.7 26.7C23.1 26.3 23.1 25.7 22.7 25.3ZM33 25H26C25.4 25 25 25.4 25 26V33C25 33.6 25.4 34 26 34H33C33.6 34 34 33.6 34 33V26C34 25.4 33.6 25 33 25ZM32 32H27V27H32V32Z" fill="#00d1b2"></path>
                    <circle cx="24" cy="24" r="23.5" stroke="#00d1b2"></circle>
                  </svg>
                </div>
                <div>
                  <h4 class="mb-2 is-size-5 has-text-weight-bold">Blocking unnecessary information</h4>
                  <p>No banners, popups, video-preview advs</p>
                </div>
              </div>
              <div class="column is-4 is-flex">
                <div class="mr-4">
                  <svg viewbox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg" style="width: 48px;height: 48px">
                    <path d="M25.6 22.9C25.7 23 25.8 23 26 23H33C33.6 23 34 22.6 34 22C34 21.8 34 21.7 33.9 21.6L30.4 14.6C30.1 14.1 29.5 13.9 29 14.2C28.9 14.3 28.7 14.4 28.6 14.6L25.1 21.6C24.9 22 25.1 22.6 25.6 22.9ZM29.5 17.2L31.4 21H27.6L29.5 17.2ZM18.5 14C16 14 14 16 14 18.5C14 21 16 23 18.5 23C21 23 23 21 23 18.5C23 16 21 14 18.5 14ZM18.5 21C17.1 21 16 19.9 16 18.5C16 17.1 17.1 16 18.5 16C19.9 16 21 17.1 21 18.5C21 19.9 19.9 21 18.5 21ZM22.7 25.3C22.3 24.9 21.7 24.9 21.3 25.3L18.5 28.1L15.7 25.3C15.3 24.9 14.7 24.9 14.3 25.3C13.9 25.7 13.9 26.3 14.3 26.7L17.1 29.5L14.3 32.3C13.9 32.7 13.9 33.3 14.3 33.7C14.7 34.1 15.3 34.1 15.7 33.7L18.5 30.9L21.3 33.7C21.7 34.1 22.3 34.1 22.7 33.7C23.1 33.3 23.1 32.7 22.7 32.3L19.9 29.5L22.7 26.7C23.1 26.3 23.1 25.7 22.7 25.3ZM33 25H26C25.4 25 25 25.4 25 26V33C25 33.6 25.4 34 26 34H33C33.6 34 34 33.6 34 33V26C34 25.4 33.6 25 33 25ZM32 32H27V27H32V32Z" fill="#00d1b2"></path>
                    <circle cx="24" cy="24" r="23.5" stroke="#00d1b2"></circle>
                  </svg>
                </div>
                <div>
                  <h4 class="mb-2 is-size-5 has-text-weight-bold">Different servers locations</h4>
                  <p>Servers located on the different countries</p>
                </div>
              </div>
            </div>
          </div>
        </section>

        <footer class="section has-background-light">
          <div class="container">
            <div class="pb-5 is-flex is-flex-wrap-wrap is-justify-content-between is-align-items-center">
              <div class="mr-auto mb-2">
                <a class="is-inline-block" href="https://lab.sys-adm.in" target="_blank">
                  <img class="image" src="https://lab.sys-adm.in/images/lab.sys-adm.in-logo.png" alt="" width="96px">
                </a>
              </div>
              <div>

                <ul class="is-flex is-flex-wrap-wrap is-align-items-center is-justify-content-center">
                {{range .Links}}
                  <li class="mr-4"><a class="button is-white" href="{{.URL}}" target="_blank">{{.Title}}</a></li>
                {{end}}
                </ul>

                <!-- <ul class="is-flex is-flex-wrap-wrap is-align-items-center is-justify-content-center">
                  <li class="mr-4"><a class="button is-white" href="#">Blocky repo</a></li>
                  <li class="mr-4"><a class="button is-white" href="#">BLD repo</a></li>
                </ul> -->
              </div>
            </div>
          </div>
          <div class="pt-5" style="border-top: 1px solid #dee2e6;"></div>
          <div class="container">
            <div class="is-flex-tablet is-justify-content-between is-align-items-center">
              <p>BLD © 2021</p>
              <div class="py-2 is-hidden-tablet"></div>
              <div class="ml-auto">
                <a class="mr-4 is-inline-block" href="#">
                  <img src="bulma-plain-assets/socials/facebook.svg" alt="">
                </a>
                <a class="mr-4 is-inline-block" href="#">
                  <img src="bulma-plain-assets/socials/github.svg" alt="">
                </a>
              </div>
            </div>
          </div>
      </footer>
    </body>

  </html>
`
