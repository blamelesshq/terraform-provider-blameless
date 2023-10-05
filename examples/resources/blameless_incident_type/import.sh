# This resource can be imported by specifying the incident type id.
#
# Example:
terraform import blameless_incident_type.incident_type_security {{incident type id}}

# To get the incident type ids call the Blameless API
URL=https://${INSTANCE}.blameless.io/api/v2/settings/incidents/types

curl --request GET \
     --url $URL    \
     --header "Authorization: Bearer $TOKEN"
     