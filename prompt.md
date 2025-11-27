CIVILINK - AI-Powered Citizen-Government Engagement Platform for Kwara State, Nigeria

Build a full-stack web application using React, Tailwind CSS, and a backend-as-a-service with the following complete specifications:

DESIGN SYSTEM
Colors:

Primary Green: #1B7339 (Nigerian government green)
Secondary Green: #006B3F
Accent Gold: #FFD700
Emergency Red: #DC2626
Success Green: #10B981
Warning Orange: #F59E0B
Info Blue: #3B82F6
Header Stripe: A thin horizontal bar with green-white-green sections representing the Nigerian flag.

Typography: Clean, modern sans-serif. Large text mode increases base font by 25%.

Accessibility Modes:

Normal mode (default)
High Contrast mode (black background, white text, yellow links)
Large Text mode (125% font size, 48px minimum touch targets)
Text-to-Speech toggle (uses browser SpeechSynthesis API)
Multi-Language Support: English (default), Yoruba, Hausa, Igbo. Store preference in localStorage. All UI text uses a translation function t('key') that returns the appropriate translation.

ENTITIES (Database Schema)
1. Report
{
  "report_number": "string (format: CL-2025-XXXXXX)",
  "category": "enum: road_infrastructure, electricity, water_sanitation, security, waste_management, flood, corruption, emergency, health, education, other",
  "title": "string",
  "description": "string",
  "location_lga": "enum: Asa, Baruten, Edu, Ekiti, Ifelodun, Ilorin East, Ilorin South, Ilorin West, Irepodun, Isin, Kaiama, Moro, Offa, Oke Ero, Oyun, Pategi",
  "location_address": "string",
  "location_lat": "number",
  "location_lng": "number",
  "status": "enum: submitted, assigned, in_progress, resolved (default: submitted)",
  "priority": "enum: low, medium, high, critical (default: medium)",
  "is_emergency": "boolean (default: false)",
  "assigned_agency": "string",
  "assigned_officer_email": "string",
  "media_urls": "array of strings",
  "is_anonymous": "boolean (default: false)",
  "public_notes": "string",
  "internal_notes": "string",
  "resolved_at": "datetime",
  "feedback_submitted": "boolean (default: false)",
  "satisfaction_score": "number 1-5",
  "feedback_text": "string",
  "feedback_sentiment": "enum: positive, neutral, negative",
  "ai_insights": "string",
  "resolution_score": "number 0-100"
}
2. Agency
{
  "name": "string",
  "category": "enum matching report categories",
  "description": "string",
  "contact_email": "email",
  "contact_phone": "string",
  "is_active": "boolean (default: true)"
}
3. Event
{
  "title": "string",
  "description": "string",
  "category": "enum: government, health, tax, infrastructure, housing, education, agriculture, cultural, employment, scholarship",
  "event_date": "string",
  "end_date": "date",
  "location": "string",
  "target_lgas": "array of strings",
  "priority": "enum: low, medium, high (default: medium)",
  "action_button_text": "string",
  "action_button_link": "string",
  "is_active": "boolean (default: true)",
  "image_url": "string"
}
4. Document
{
  "title": "string",
  "category": "enum: work, personal_id, certificates, medical, property, family, generated",
  "file_url": "string",
  "file_type": "string",
  "file_size_kb": "number",
  "expiry_date": "date",
  "tags": "array of strings",
  "generated_content": "string",
  "document_type": "string"
}
5. Notification
{
  "user_email": "string",
  "title": "string",
  "message": "string",
  "type": "enum: report_update, event_reminder, document_expiry, system",
  "link": "string",
  "is_read": "boolean (default: false)"
}
6. AIInsight
{
  "insight_type": "enum: pattern, hotspot, intervention, trend, alert",
  "title": "string",
  "description": "string",
  "affected_lgas": "array of strings",
  "affected_categories": "array of strings",
  "severity": "enum: low, medium, high, critical (default: medium)",
  "recommended_actions": "array of strings",
  "report_count": "number",
  "is_active": "boolean (default: true)",
  "expires_at": "datetime"
}
7. OfficerMetrics
{
  "officer_email": "string",
  "officer_name": "string",
  "agency_name": "string",
  "total_assigned": "number (default: 0)",
  "total_resolved": "number (default: 0)",
  "avg_resolution_hours": "number",
  "avg_satisfaction_score": "number",
  "resolution_rate": "number",
  "performance_score": "number 0-100",
  "last_updated": "datetime"
}
User Entity (Extended)
Add these fields to the built-in User entity:

{
  "phone": "string",
  "lga": "string",
  "is_officer": "boolean (default: false)",
  "agency_name": "string",
  "ministry": "string",
  "grade_level": "number",
  "employment_date": "date",
  "last_promotion_date": "date"
}
PAGES
1. Home Page
Hero section with gradient background (primary to secondary green)
Title: "CIVILINK - KWARA STATE"
Subtitle: "Report Issues. Track Progress. Stay Informed."
Two CTA buttons: "Submit a Report" and "View Events"
Four feature badges: AI-Powered Routing, Real-Time Tracking, Multi-Language, Accessible to All
Statistics section showing total reports, resolved, in progress from database
Preview of 3 upcoming events in card grid
Features grid (6 cards): Report Issues, Track Progress, Events & Announcements, AI Document Assistant, Document Vault, Issue Map
Quick action links section
"Join Today" CTA for non-authenticated users
2. Dashboard (Authenticated)
Welcome message with user's name
4 stat cards: My Reports count, My Documents count, Upcoming Events count, Unread Notifications count
"What's Happening in Kwara" section - 3 most recent active events
AI Document Assistant CTA card with two buttons: Check Promotion Eligibility, Generate Leave Letter
"Your Reports" section with tabs for each status (All, Submitted, Assigned, In Progress, Resolved)
Each report shown as a card with status badge, priority, category icon, LGA, date
3. Submit Report Page
Form with fields:
Title (required)
Description (required, with AI classification trigger on blur)
Category dropdown (auto-selected by AI based on description)
LGA dropdown (16 Kwara State LGAs)
Specific Address (optional)
Photo upload (multiple, max 4, stored via UploadFile integration)
"This is an Emergency" toggle (shows red warning when enabled)
"Submit Anonymously" toggle
AI Classification: When description loses focus and has 20+ chars, call LLM to suggest category and priority
On submit: Generate report_number (CL-2025-XXXXXX format), auto-assign agency based on category, create notification
Success view: Shows tracking number, options to go to dashboard or submit another
4. Events Page
Header with title and description
Search input to filter by title/description
Category tabs for filtering (All + 10 categories)
Event cards showing: category badge, title, description, date, location, priority indicator, share button
Loading skeleton cards while fetching
5. Issue Map Page
Interactive map using React-Leaflet centered on Kwara State (lat: 8.9, lng: 4.5, zoom: 8)
Sidebar with:
Category filter dropdown
Status filter dropdown
Legend showing category colors
Emergency reports list (if any)
Custom markers colored by category
Click marker to show report details popup
Selected report details panel with full information
LGA Coordinates (hardcoded):

{
  "Ilorin West": [8.4966, 4.5426],
  "Ilorin East": [8.5000, 4.6000],
  "Ilorin South": [8.4500, 4.5500],
  // ... all 16 LGAs
}
6. Documents Page (Authenticated)
Storage usage indicator (mock: 2.4GB of 5GB)
Search input and category filter tabs
Upload button opening modal with:
Title input
Category dropdown
Expiry date (optional)
File input (accepts pdf, doc, docx, png, jpg, jpeg)
Document grid with cards showing:
File type icon (color-coded)
Title
Category badge
Expiry warning badge if within 30 days or expired
View/Download and Delete buttons
7. AI Assistant Page (Authenticated)
Two tabs:

Tab 1: Promotion Eligibility Checker

Grade Level dropdown (1-17)
Last Promotion Date input
Employment Date input
"Check Eligibility" button
Result card showing:
Eligible/Not Eligible status with icon
Current and Next grade levels
Years at current level
Eligibility date
If eligible: "Draft Application Letter Now" button
If not eligible: "Set Reminder" button
Tab 2: Document Generator

Document type selection (6 buttons):
Promotion Application
Leave Application
Transfer Request
Memo to Supervisor
Job Application
Complaint Letter
Common fields: Full Name, Ministry/Agency
Conditional fields based on document type:
Leave: Leave type dropdown, start/end dates, supervisor name, reason
Transfer: Current location, desired location, reason
Memo/Job/Complaint: Subject, key points
"Generate Document" button calls LLM with detailed prompt
Generated document preview with copy and download buttons
Auto-save generated document to Documents entity
8. Submit Feedback Page
Accessed via ?reportId=xxx parameter
Shows report title and tracking number
5-star rating selector with hover effects
Rating labels: Very Dissatisfied, Dissatisfied, Neutral, Satisfied, Very Satisfied
Optional text feedback textarea
On submit:
Call LLM to analyze sentiment (returns: positive/neutral/negative)
Update report with satisfaction_score, feedback_text, feedback_sentiment, feedback_submitted=true
Success view with thank you message
9. Settings Page (Authenticated)
Basic Information section: Full Name (readonly), Email (readonly), Phone, LGA dropdown
Officer Settings section:
"I am a Government Officer" toggle
Agency dropdown (shown when toggle is on)
Civil Service Details section:
Ministry/Agency input
Grade Level dropdown (1-17)
Employment Date
Last Promotion Date
Testing Guide card explaining how to access Officer Portal (toggle on + select agency) and Admin Portal (requires admin role)
Save button to update user profile
10. Officer Portal (Officers/Admins only)
Access check: user.is_officer must be true OR user.role must be 'admin'
Header with agency name and officer name
Performance Card (left column):

Large performance score (0-100)
Progress bar
4 metric boxes: Resolution Rate %, Avg Time (hours), Satisfaction rating, Resolved/Total count
Performance badge: Excellent (80+), Good (60+), Fair (40+), Needs Improvement (<40)
Stats row (right column): Total, Pending, In Progress, Resolved, Emergency (with flash animation if >0)

Filters card: Search input, Status dropdown, LGA dropdown, Refresh button, Export CSV button

Reports table:

Columns: Report #, Title, Category, LGA, Status, Priority, Date, Action
Emergency reports highlighted with red background and flash icon
Click "View" opens dialog with:
Report details (title, description, location, date, photos)
Status update dropdown
Public Note textarea (visible to citizen)
Internal Note textarea (officers only)
Update button
On status update:

If resolved: Set resolved_at, calculate resolution_score (base 50 + 25 for good public note + 15 for internal note + 10 for proper workflow)
Create notification to citizen
If resolved: Notification includes link to SubmitFeedback page
11. Admin Portal (Admins only)
Access check: user.role must be 'admin'
Stats row (7 cards): Total Reports, Pending, In Progress, Resolved, Emergency, Total Agencies, Active Events

Performance metrics row (3 gradient cards):

Resolution Rate % (green gradient)
Avg Resolution Time in days (blue gradient)
Reports This Month (purple gradient)
Charts section (2x2 grid):

Reports Over Time - Line chart with date range selector (7/14/30 days)
Status Distribution - Donut/pie chart with 4 segments
Reports by Category - Horizontal bar chart
Top LGAs by Reports - Vertical bar chart (top 10)
AI Insights Panel:

"Generate Insights" button calls LLM with aggregated report data
LLM prompt asks for:
Pattern detection (recurring issues)
Geographic hotspots
Recommended government interventions
Trend analysis
Critical alerts
Response JSON schema with array of insights
Each insight displayed as card with:
Icon based on type (pattern=TrendingUp, hotspot=MapPin, intervention=Lightbulb, alert=AlertTriangle)
Severity badge (color-coded)
Title and description
Affected LGAs badges
Recommended Actions in green box with bullet points
Report count and generation timestamp
Satisfaction Metrics Card:

Average star rating with visual stars
Sentiment breakdown: 3 boxes showing Positive %, Neutral %, Negative %
Rating distribution: 5 progress bars for each star level with counts
Performance Metrics (2 cards side by side):

Agency Performance:

Top 5 agencies ranked by performance score
Each shows: name, score badge, reports count, resolution %, avg time, satisfaction
Top Officers:

Top 5 officers with rank medals (gold/silver/bronze for 1-2-3)
Each shows: email username, agency, score badge, resolved/total
Emergency Reports Section: (only if emergency count > 0)

Red border card with flash animation
List of active emergency reports with details
LAYOUT COMPONENT
Wraps all pages with Header and Footer
Provides LanguageProvider and AccessibilityProvider contexts
Injects CSS variables and accessibility mode classes
CSS for emergency-flash animation, high-contrast mode overrides, large-text mode adjustments
HEADER COMPONENT
Logo: Circle with "CL" text + "CIVILINK" title + "Your Government, In Your Pocket" subtitle
Desktop nav: Home, Events, Issue Map, My Documents + Dashboard (if authenticated) + Officer Portal (if is_officer) + Admin Portal (if admin)
Right side actions:
Language selector dropdown with flag emojis
Accessibility dropdown with mode toggles
Notifications bell with unread count badge
User avatar dropdown with profile/settings/logout
Sign In button (if not authenticated)
Mobile: Hamburger menu opening Sheet with all navigation + language/accessibility options
FOOTER COMPONENT
Partners section: List of 4 key partners with descriptions
Brand section: Logo, tagline, description
Quick Links: Home, Events, Issue Map, Documents
Services: Submit Report, AI Document Assistant, Document Vault
Contact: Address (Kwara State Secretariat, Ilorin), Email
Copyright bar with Kwara AI Summit link
CONTEXT PROVIDERS
LanguageContext
State: language (en/yo/ha/ig)
Persist to localStorage key 'civilink_language'
Provides: language, setLanguage, t(key) function
Full translations object for all 4 languages
AccessibilityContext
State: mode (normal/high_contrast/large_text), textToSpeech boolean
Persist to localStorage keys 'civilink_accessibility_mode' and 'civilink_tts'
Provides: mode, setMode, textToSpeech, setTextToSpeech, speak(text), stopSpeaking(), getAccessibilityClasses()
KEY INTERACTIONS
Report Submission Flow:

User fills form → AI classifies description → Auto-assigns category/priority
On submit → Generate unique report number → Match agency by category → Create report → Create notification → Show success
Report Resolution Flow:

Officer views report → Updates status to resolved → System calculates resolution_score → Creates notification with feedback link
Citizen clicks notification → Rates 1-5 stars + optional comment → AI analyzes sentiment → Updates report
AI Insights Generation:

Admin clicks "Generate Insights" → System aggregates all report data by LGA and category
LLM analyzes patterns and returns structured insights → Old insights deactivated → New insights saved and displayed
Performance Score Calculation:

Resolution Rate: 40% weight (resolved/total)
Speed: 30% weight (based on avg resolution hours, benchmark 168 hours = 1 week)
Satisfaction: 30% weight (avg satisfaction / 5)
DEMO DATA TO PRE-POPULATE
Create sample agencies for each category, sample events across categories, and sample reports with various statuses, priorities, and LGAs to demonstrate the platform's capabilities.

This prompt should enable an AI builder to recreate the complete CIVILINK platform with all features, styling, and functionality intact.
