# nf_nwdaf
Network Data Analytics Function - my5GC

To perform this step, do the following:
1. Build nwdaf image ``` docker build -f DockerFile -t ciromacedo/nf_nwdaf . ```
2. Run nwdaf image ```docker run -it -d -p 5003:5003 --name nf_nwdaf ciromacedo/nf_nwdaf bash ```
3. Start nwdaf server ```docker exec -it -d nf_nwdaf /bin/sh" ```
4. Publish DockerHub ```docker push ciromacedo/nf_nwdaf ```