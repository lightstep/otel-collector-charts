
{{- define "kube-otel-stack.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 50 | trimSuffix "-" -}}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
The components in this chart create additional resources that expand the longest created name strings.
The longest name that gets created adds and extra 37 characters, so truncation should be 63-35=26.
*/}}
{{- define "kube-otel-stack.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 26 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- .Release.Name | trunc 26 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 26 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{/*
Allow the release namespace to be overridden for multi-namespace deployments in combined charts
*/}}
{{- define "kube-otel-stack.namespace" -}}
  {{- if .Values.namespaceOverride -}}
    {{- .Values.namespaceOverride -}}
  {{- else -}}
    {{- .Release.Namespace -}}
  {{- end -}}
{{- end -}}

{{/* Create chart name and version as used by the chart label. */}}
{{- define "kube-otel-stack.chartref" -}}
{{- replace "+" "_" .Chart.Version | printf "%s-%s" .Chart.Name -}}
{{- end }}

{{/* Generate basic labels */}}
{{- define "kube-otel-stack.labels" }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/version: "{{ replace "+" "_" .Chart.Version }}"
app.kubernetes.io/part-of: {{ template "kube-otel-stack.name" . }}
chart: {{ template "kube-otel-stack.chartref" . }}
release: {{ $.Release.Name | quote }}
heritage: {{ $.Release.Service | quote }}
{{- if .Values.commonLabels}}
{{ toYaml .Values.commonLabels }}
{{- end }}
{{- end }}

{{/* Allow KubeVersion to be overridden. */}}
{{- define "kube-otel-stack.kubeVersion" -}}
  {{- default .Capabilities.KubeVersion.Version .Values.kubeVersionOverride -}}
{{- end -}}

{{/* Get value based on current Kubernetes version */}}
{{- define "kube-otel-stack.kubeVersionDefaultValue" -}}
  {{- $values := index . 0 -}}
  {{- $kubeVersion := index . 1 -}}
  {{- $old := index . 2 -}}
  {{- $new := index . 3 -}}
  {{- $default := index . 4 -}}
  {{- if kindIs "invalid" $default -}}
    {{- if semverCompare $kubeVersion (include "kube-otel-stack.kubeVersion" $values) -}}
      {{- print $new -}}
    {{- else -}}
      {{- print $old -}}
    {{- end -}}
  {{- else -}}
    {{- print $default }}
  {{- end -}}
{{- end -}}

{{/* Get value for kube-controller-manager depending on insecure scraping availability */}}
{{- define "kube-otel-stack.kubeControllerManager.insecureScrape" -}}
  {{- $values := index . 0 -}}
  {{- $insecure := index . 1 -}}
  {{- $secure := index . 2 -}}
  {{- $userValue := index . 3 -}}
  {{- include "kube-otel-stack.kubeVersionDefaultValue" (list $values ">= 1.22-0" $insecure $secure $userValue) -}}
{{- end -}}

{{/* Get value for kube-scheduler depending on insecure scraping availability */}}
{{- define "kube-otel-stack.kubeScheduler.insecureScrape" -}}
  {{- $values := index . 0 -}}
  {{- $insecure := index . 1 -}}
  {{- $secure := index . 2 -}}
  {{- $userValue := index . 3 -}}
  {{- include "kube-otel-stack.kubeVersionDefaultValue" (list $values ">= 1.23-0" $insecure $secure $userValue) -}}
{{- end -}}
