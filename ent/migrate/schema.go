// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AppointmentsColumns holds the columns for the "appointments" table.
	AppointmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "status", Type: field.TypeString},
		{Name: "period", Type: field.TypeJSON},
		{Name: "patient_appointments", Type: field.TypeUUID, Nullable: true},
		{Name: "provider_appointments", Type: field.TypeUUID, Nullable: true},
	}
	// AppointmentsTable holds the schema information for the "appointments" table.
	AppointmentsTable = &schema.Table{
		Name:       "appointments",
		Columns:    AppointmentsColumns,
		PrimaryKey: []*schema.Column{AppointmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "appointments_patients_appointments",
				Columns:    []*schema.Column{AppointmentsColumns[3]},
				RefColumns: []*schema.Column{PatientsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "appointments_providers_appointments",
				Columns:    []*schema.Column{AppointmentsColumns[4]},
				RefColumns: []*schema.Column{ProvidersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DiagnosesColumns holds the columns for the "diagnoses" table.
	DiagnosesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "status", Type: field.TypeString},
		{Name: "last_updated", Type: field.TypeTime},
		{Name: "code", Type: field.TypeJSON},
	}
	// DiagnosesTable holds the schema information for the "diagnoses" table.
	DiagnosesTable = &schema.Table{
		Name:       "diagnoses",
		Columns:    DiagnosesColumns,
		PrimaryKey: []*schema.Column{DiagnosesColumns[0]},
	}
	// FeedbacksColumns holds the columns for the "feedbacks" table.
	FeedbacksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "status", Type: field.TypeString},
		{Name: "start_time", Type: field.TypeTime},
		{Name: "completion_time", Type: field.TypeTime},
		{Name: "patient_feedbacks", Type: field.TypeUUID, Nullable: true},
		{Name: "survey_feedbacks", Type: field.TypeUUID, Nullable: true},
	}
	// FeedbacksTable holds the schema information for the "feedbacks" table.
	FeedbacksTable = &schema.Table{
		Name:       "feedbacks",
		Columns:    FeedbacksColumns,
		PrimaryKey: []*schema.Column{FeedbacksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "feedbacks_patients_feedbacks",
				Columns:    []*schema.Column{FeedbacksColumns[4]},
				RefColumns: []*schema.Column{PatientsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "feedbacks_surveys_feedbacks",
				Columns:    []*schema.Column{FeedbacksColumns[5]},
				RefColumns: []*schema.Column{SurveysColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PatientsColumns holds the columns for the "patients" table.
	PatientsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeJSON},
		{Name: "gender", Type: field.TypeString},
		{Name: "birtdate", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"postgres": "date"}},
		{Name: "contact", Type: field.TypeJSON, Nullable: true},
		{Name: "address", Type: field.TypeJSON, Nullable: true},
	}
	// PatientsTable holds the schema information for the "patients" table.
	PatientsTable = &schema.Table{
		Name:       "patients",
		Columns:    PatientsColumns,
		PrimaryKey: []*schema.Column{PatientsColumns[0]},
	}
	// PromptsColumns holds the columns for the "prompts" table.
	PromptsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "sort_order", Type: field.TypeInt},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "response_type", Type: field.TypeJSON},
		{Name: "additional_feedback", Type: field.TypeBool, Default: false},
		{Name: "survey_prompts", Type: field.TypeUUID, Nullable: true},
	}
	// PromptsTable holds the schema information for the "prompts" table.
	PromptsTable = &schema.Table{
		Name:       "prompts",
		Columns:    PromptsColumns,
		PrimaryKey: []*schema.Column{PromptsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "prompts_surveys_prompts",
				Columns:    []*schema.Column{PromptsColumns[6]},
				RefColumns: []*schema.Column{SurveysColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PromptResponsesColumns holds the columns for the "prompt_responses" table.
	PromptResponsesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "prompt_index", Type: field.TypeInt},
		{Name: "range_value", Type: field.TypeInt, Nullable: true},
		{Name: "bool_value", Type: field.TypeString, Nullable: true},
		{Name: "enum_value", Type: field.TypeJSON, Nullable: true},
		{Name: "label_values", Type: field.TypeJSON, Nullable: true},
		{Name: "freeform_value", Type: field.TypeString, Nullable: true},
		{Name: "feedback_responses", Type: field.TypeUUID, Nullable: true},
	}
	// PromptResponsesTable holds the schema information for the "prompt_responses" table.
	PromptResponsesTable = &schema.Table{
		Name:       "prompt_responses",
		Columns:    PromptResponsesColumns,
		PrimaryKey: []*schema.Column{PromptResponsesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "prompt_responses_feedbacks_responses",
				Columns:    []*schema.Column{PromptResponsesColumns[7]},
				RefColumns: []*schema.Column{FeedbacksColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProvidersColumns holds the columns for the "providers" table.
	ProvidersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeJSON},
	}
	// ProvidersTable holds the schema information for the "providers" table.
	ProvidersTable = &schema.Table{
		Name:       "providers",
		Columns:    ProvidersColumns,
		PrimaryKey: []*schema.Column{ProvidersColumns[0]},
	}
	// SurveysColumns holds the columns for the "surveys" table.
	SurveysColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"unpublished", "active", "archived"}, Default: "unpublished"},
		{Name: "active_time", Type: field.TypeTime, Nullable: true},
		{Name: "archive_time", Type: field.TypeTime, Nullable: true},
	}
	// SurveysTable holds the schema information for the "surveys" table.
	SurveysTable = &schema.Table{
		Name:       "surveys",
		Columns:    SurveysColumns,
		PrimaryKey: []*schema.Column{SurveysColumns[0]},
	}
	// AppointmentDiagnosesColumns holds the columns for the "appointment_diagnoses" table.
	AppointmentDiagnosesColumns = []*schema.Column{
		{Name: "appointment_id", Type: field.TypeUUID},
		{Name: "diagnosis_id", Type: field.TypeUUID},
	}
	// AppointmentDiagnosesTable holds the schema information for the "appointment_diagnoses" table.
	AppointmentDiagnosesTable = &schema.Table{
		Name:       "appointment_diagnoses",
		Columns:    AppointmentDiagnosesColumns,
		PrimaryKey: []*schema.Column{AppointmentDiagnosesColumns[0], AppointmentDiagnosesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "appointment_diagnoses_appointment_id",
				Columns:    []*schema.Column{AppointmentDiagnosesColumns[0]},
				RefColumns: []*schema.Column{AppointmentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "appointment_diagnoses_diagnosis_id",
				Columns:    []*schema.Column{AppointmentDiagnosesColumns[1]},
				RefColumns: []*schema.Column{DiagnosesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AppointmentsTable,
		DiagnosesTable,
		FeedbacksTable,
		PatientsTable,
		PromptsTable,
		PromptResponsesTable,
		ProvidersTable,
		SurveysTable,
		AppointmentDiagnosesTable,
	}
)

func init() {
	AppointmentsTable.ForeignKeys[0].RefTable = PatientsTable
	AppointmentsTable.ForeignKeys[1].RefTable = ProvidersTable
	FeedbacksTable.ForeignKeys[0].RefTable = PatientsTable
	FeedbacksTable.ForeignKeys[1].RefTable = SurveysTable
	PromptsTable.ForeignKeys[0].RefTable = SurveysTable
	PromptResponsesTable.ForeignKeys[0].RefTable = FeedbacksTable
	AppointmentDiagnosesTable.ForeignKeys[0].RefTable = AppointmentsTable
	AppointmentDiagnosesTable.ForeignKeys[1].RefTable = DiagnosesTable
}
