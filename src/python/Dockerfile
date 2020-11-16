FROM python:3.7-slim as base

FROM base as builder

RUN pip install --upgrade pip
# get packages
COPY requirements.txt .

RUN pip install -r requirements.txt

FROM base as final

WORKDIR /pythonserver

# Grab packages from builder
COPY --from=builder /usr/local/lib/python3.7/ /usr/local/lib/python3.7/

# Add the application
COPY . .
ENV PORT 8082
EXPOSE 8080
ENTRYPOINT [ "python", "pythonserver.py" ]
