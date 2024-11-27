FROM gitpod/workspace-full:latest

RUN go install github.com/pulumi/upgrade-provider@main

RUN brew install gh

RUN brew install pulumi/tap/pulumi

RUN brew install pulumi/tap/pulumictl 

# Install dotnet CLI

#ARG DOTNET_VERSION="6.0"
ARG DOTNET_VERSION="8.0"

RUN curl -fsSL https://dot.net/v1/dotnet-install.sh | bash -s -- -channel ${DOTNET_VERSION}

ENV PATH "/home/gitpod/.dotnet:${PATH}"
ENV DOTNET_ROOT "/home/gitpod/.dotnet"
ENV DOTNET_SYSTEM_GLOBALIZATION_INVARIANT 1
# Allow newer dotnet version (e.g. 6) to build projects targeting older frameworks (v3.1)
ENV DOTNET_ROLL_FORWARD Major

# Install Java
RUN bash -c ". /home/gitpod/.sdkman/bin/sdkman-init.sh && \
    sdk install java 21.0.4-tem && \
    sdk default java 21.0.4-tem"

# install python?

# install JS?

#FROM --platform=linux/amd64 pulumi/pulumi:3.74.0

# Install pulumictl and set to PATH
#RUN curl -fsSL https://get.pulumi.com | sh
#ENV PATH="/root/.pulumi/bin:${PATH}"

# create a directory for pulumictl and download the binary to it and set to PATH
#RUN mkdir -p /root/pulumictl && cd /root/pulumictl/
#RUN wget https://github.com/pulumi/pulumictl/releases/download/v0.0.42/pulumictl-v0.0.42-linux-amd64.tar.gz -O /root/pulumictl/pulumictl-v0.0.42-linux-amd64.tar.gz
#RUN tar -xvf /root/pulumictl/pulumictl-v0.0.42-linux-amd64.tar.gz -C /root/pulumictl/
#ENV PATH="//root/pulumictl/:${PATH}"

#RUN apt install sudo -y

#RUN type -p curl >/dev/null || (sudo apt update && sudo apt install curl -y)
#RUN curl -fsSL https://cli.github.com/packages/githubcli-archive-keyring.gpg | sudo dd of=/usr/share/keyrings/githubcli-archive-keyring.gpg
#RUN sudo chmod go+r /usr/share/keyrings/githubcli-archive-keyring.gpg
#RUN echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main" | sudo tee /etc/apt/sources.list.d/github-cli.list > /dev/null
#RUN sudo apt update
#RUN sudo apt install gh -y

#RUN go install github.com/pulumi/upgrade-provider@main
