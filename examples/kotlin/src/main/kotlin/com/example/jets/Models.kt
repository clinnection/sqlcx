// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0

package com.example.jets

data class Jet (
  val id: Int,
  val pilotId: Int,
  val age: Int,
  val name: String,
  val color: String
)

data class Language (
  val id: Int,
  val language: String
)

data class Pilot (
  val id: Int,
  val name: String
)

data class PilotLanguage (
  val pilotId: Int,
  val languageId: Int
)

