<h1>Firebase Manager</h1>
<p>Handling many tokens from  firebase in programs, servers, applications that have many users is common. This project provide a simple server to manage your users tokens and their score and send notifications by simple requests</p>
<p>You can set the token api in your application to send its token every time and send notifications by other apis, additionally you my change your users scores to send specific notifications to specific users.</p>

<h3>Installing</h3>
<p>It is a simple docker.</p>

<h3>API Doc</h3>
<p><b>http://localhost:4300/</b> follows:</p>
<p><b>GET: /add_user/{username}/{score}</b> add new user to users list </p>
<p><b>GET: /remove_user/{username}</b> remove username from users list</p>
<p><b>GET: /set_token/{username}/{token}</b> This api should place in application to periodically send token and username, for example in react native you put this api after notifications.getToken()</p>
<p><b>POST /add_multiple_score</b></p>
