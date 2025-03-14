# TSDR-003 Docker

## Status

Accepted

## Context

Which docker image to base our postgres image on. 

## Decision

[Bitnami PostgreSQL Image](https://hub.docker.com/r/bitnami/postgresql)

## Why / Notes

- It is just a matter of time before we need the additional features. 

A google search: "what is the difference between PostreSQL Docker images binami and postgress"

The primary difference between Bitnami's PostgreSQL Docker image and the official PostgreSQL Docker image lies in added features, security enhancements, and ease of use provided by Bitnami, particularly geared towards production environments and Kubernetes deployments. 
Bitnami PostgreSQL Image:

Non-root by default:
Bitnami's image runs as a non-root user, enhancing security by limiting potential damage from vulnerabilities within the container. This is a recommended practice for production deployments. 

Replication support:
Bitnami's image supports replication environment variables, allowing you to easily configure and deploy PostgreSQL in a replicated setup. The official image does not natively support replication. 
Simplified configuration:
Bitnami's image offers a streamlined configuration experience, particularly for Kubernetes deployments, simplifying the process of setting up and managing PostgreSQL clusters. 
Enhanced security features:
Bitnami's image is designed with security best practices in mind, including non-root user execution and support for Kubernetes security constraints like those found in OpenShift. 
Helm charts for Kubernetes:
Bitnami provides Helm charts to simplify deploying and managing PostgreSQL in Kubernetes environments. 
Focus on production readiness:
Bitnami's image is tailored for production environments, offering features like replication and security enhancements to ensure reliable and robust database deployments. 

Official PostgreSQL Docker Image:

Minimalist approach:
The official image provides a bare bones PostgreSQL installation, focusing on simplicity and basic functionality. 

No replication support:
The official image does not natively support replication environment variables, making it less suitable for complex production setups requiring replication. 

Root user by default:
The official image runs as root by default, potentially increasing security risks in production environments. 

Limited configuration options:
The official image offers fewer configuration options compared to Bitnami's, particularly when it comes to security and replication. 

Designed for development and testing:
The official image is often used for development, testing, or simple deployments where replication and advanced security features are not critical. 

In essence, the Bitnami PostgreSQL Docker image provides a more comprehensive and production-ready solution with enhanced security, replication support, and simplified Kubernetes integration, making it a preferred choice for deploying PostgreSQL in complex or production environments. 

## Consequences

## Other Options

Possibilities:
- [Official PostgreSQL Image](https://hub.docker.com/_/postgres)