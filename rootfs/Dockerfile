FROM quay.io/deis/base:v0.3.5

# Add user and group
RUN adduser --system \
	--shell /bin/bash \
	--disabled-password \
	--home /opt/helm-ui \
	--group \
	helm-ui

COPY . /

# Fix some permission since we'll be running as a non-root user
RUN mkdir -p /opt/helm-ui/repository/cache \
    && chown -R helm-ui:helm-ui /opt/helm-ui \
    && chmod +x /opt/helm-ui/sbin/* 

USER helm-ui
WORKDIR /opt/helm-ui/sbin/
CMD ["./helm-ui"]

ARG VERSION
ARG BUILD_DATE
ENV VERSION $VERSION
ENV BUILD_DATE $BUILD_DATE
ENV PATH $PATH:/opt/discapi/sbin