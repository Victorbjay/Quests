# CiviLink MVP - Complete Testing Guide (v2.0)

## Overview
CiviLink is a verified citizen reporting system for Kwara State, Nigeria with multi-language support, accessibility features, and mobile-responsive design.

---

## 🎯 New Features Added (January 2025)

### ✨ What's New
1. **🌍 Multi-Language Support**: English, Yoruba, Hausa, Igbo
2. **♿ Accessibility Features**: Text-to-Speech, High Contrast, Large Text
3. **📊 Public Statistics Dashboard**: Real-time transparent data visualization
4. **🗺️ Interactive Map View**: Visual report locations across Kwara State
5. **📱 Mobile Hamburger Menu**: Responsive navigation for all devices
6. **🏛️ Official Government Branding**: Nigerian Coat of Arms & Kwara State Emblem
7. **🏢 Real Kwara State Agencies**: All 10+ government agencies with actual contact info

---

## 📋 Quick Access URLs

- **Landing Page**: `/`
- **Public Statistics**: `/statistics`
- **Interactive Map**: `/map`
- **Submit Report**: `/submit-report`
- **Authentication**: `/auth`
- **Citizen Dashboard**: `/dashboard`
- **Officer Portal**: `/officer`
- **Admin Portal**: `/admin`

---

## 🔧 Pre-Testing Setup

### 1. Create Test Accounts

Sign up for 3 different user accounts:

#### **Account 1: Citizen User**
- Email: `citizen@test.com`
- Password: `test123456`
- Role: Citizen (default)

#### **Account 2: Officer User**
- Email: `officer@test.com`
- Password: `test123456`
- Role: Officer (needs manual upgrade - see instructions below)

#### **Account 3: Admin User**
- Email: `admin@test.com`
- Password: `test123456`
- Role: Admin (needs manual upgrade - see instructions below)

### 2. Upgrade User Roles

After signing up, upgrade roles in the database:

**Via Lovable Cloud UI**:
1. Click "Cloud" in top navigation
2. Go to Database → Tables → `profiles`
3. Find each test user and edit:
   - Set `officer@test.com` role to `officer` and assign an `agency_id`
   - Set `admin@test.com` role to `admin`
   ```

**Via SQL** (Advanced):
```sql
-- Upgrade officer
UPDATE profiles SET role = 'officer', agency_id = 'agency-uuid-here' 
WHERE email = 'officer@test.com';

-- Upgrade admin
UPDATE profiles SET role = 'admin' 
WHERE email = 'admin@test.com';
```

---

## 🆕 New Features Testing

### 🌍 Multi-Language Support

**Available Languages**: English, Yoruba, Hausa, Igbo

**Test Steps**:
1. Look for globe icon (🌐) with flag in header
2. Click to open language dropdown
3. Select each language:
   - [ ] English (🇬🇧)
   - [ ] Yoruba (🇳🇬) 
   - [ ] Hausa (🇳🇬)
   - [ ] Igbo (🇳🇬)
4. Verify translations on:
   - [ ] Navigation menu
   - [ ] Hero section
   - [ ] Form labels
   - [ ] Button text
   - [ ] Statistics page
   - [ ] Map page

**Expected**: Language persists across pages, all UI elements translate properly.

---

### ♿ Accessibility Features

#### Test Text-to-Speech
1. Click accessibility icon in header
2. Enable "Text to Speech"
3. Navigate through pages
4. Verify audio announcements for:
   - [ ] Page changes
   - [ ] Button clicks
   - [ ] Form submissions

#### Test High Contrast Mode
1. Enable "High Contrast" in accessibility menu
2. Verify:
   - [ ] Text is clearly visible
   - [ ] Strong border contrast
   - [ ] Works in light mode
   - [ ] Works in dark mode

#### Test Large Text Mode
1. Enable "Large Text"
2. Check:
   - [ ] All text sizes increase
   - [ ] Headings scale properly
   - [ ] Buttons remain usable
   - [ ] Layouts don't break

**Expected**: Settings persist in localStorage, work individually or combined.

---

### 📊 Public Statistics Dashboard (`/statistics`)

**Test Steps**:
1. Navigate to `/statistics` (no login required)
2. Verify 4 metric cards:
   - [ ] Total Reports (number)
   - [ ] Resolved Reports (green number)
   - [ ] Avg Resolution Time (days)
   - [ ] Active Agencies (10+)
3. Check Pie Chart (Reports by Category):
   - [ ] Shows all 11 categories
   - [ ] Colors match green theme
   - [ ] Interactive hover
4. Check Bar Chart (Reports by Status):
   - [ ] Shows 5 status types
   - [ ] Green bars
   - [ ] Proper scaling
5. Check Line Chart (Resolution Trends):
   - [ ] Shows last 7 days
   - [ ] Two lines (submitted/resolved)
   - [ ] Proper legend

**Mobile Test**:
- [ ] All charts responsive
- [ ] Cards stack vertically
- [ ] Charts remain readable

---

### 🗺️ Interactive Map View (`/map`)

**Test Steps**:
1. Navigate to `/map` (no login required)
2. Verify map loads centered on Kwara State
3. Test Category Filter:
   - [ ] Select "Road Infrastructure"
   - [ ] Only road reports show
   - [ ] Select "All Categories"
4. Test Status Filter:
   - [ ] Select "Resolved"
   - [ ] Only resolved reports show
   - [ ] Colors match status:
     - Orange = Submitted
     - Blue = Assigned
     - Yellow = In Progress
     - Green = Resolved
     - Gray = Closed
5. Test Combined Filters:
   - [ ] Apply both filters
   - [ ] Verify intersection works
6. Click on markers:
   - [ ] Popup shows report details
   - [ ] Title, category, status visible
   - [ ] Location address shown
   - [ ] Creation date displayed

**Mobile Test**:
- [ ] Map controls accessible
- [ ] Filters work on mobile
- [ ] Pan/zoom gestures work

---

### 📱 Mobile Responsive Navigation

#### Desktop View (≥1024px)
1. View site on desktop width
2. Verify:
   - [ ] Full navigation bar visible
   - [ ] All links horizontal
   - [ ] Emblems displayed
   - [ ] Clean, organized layout

#### Mobile View (<1024px)
1. Resize to mobile width
2. Check:
   - [ ] Hamburger icon appears (☰)
   - [ ] Desktop nav hidden
   - [ ] Emblems still visible
3. Click hamburger menu:
   - [ ] Side sheet opens from right
   - [ ] Smooth animation
4. Verify menu contains:
   - [ ] Home link
   - [ ] Statistics link
   - [ ] Map link
   - [ ] Submit Report button (full width)
   - [ ] Sign In button (full width)
   - [ ] Language switcher
   - [ ] Accessibility controls
5. Test menu behavior:
   - [ ] Clicking link closes menu
   - [ ] Clicking outside closes menu
   - [ ] X icon closes menu
6. Test at exact breakpoint (1024px):
   - [ ] Hamburger appears/disappears correctly

---

### 🏛️ Government Branding

**Verify**:
1. Header displays:
   - [ ] Nigerian Coat of Arms (left)
   - [ ] Kwara State Emblem (center-left)
   - [ ] "CiviLink" text
   - [ ] "Kwara State" subtitle
2. Color scheme:
   - [ ] Primary green is #1B7339
   - [ ] Matches Nigerian flag
   - [ ] Professional appearance
3. Hero section:
   - [ ] Background image loads
   - [ ] Shows Kwara State landscape
   - [ ] Citizens visible in image
   - [ ] Government buildings shown

---

### 🏢 Real Agency Data

**Test in Admin Portal**:
1. Login as admin
2. Go to "Agencies" tab
3. Verify all agencies present:
   - [ ] Kwara State Ministry of Works and Transport
   - [ ] Kwara State Water Corporation
   - [ ] IBEDC - Kwara Region
   - [ ] NSCDC - Kwara Command
   - [ ] Ministry of Health (Hon. Amina Ahmed El-Imam)
   - [ ] Ministry of Education
   - [ ] KWEPA (2 entries for waste & environmental)
   - [ ] ICPC - Kwara Office
   - [ ] NEMA - Kwara Operations
   - [ ] Ministry of Finance
4. Check contact details:
   - [ ] Real email addresses
   - [ ] Real phone numbers (Nigerian format)
   - [ ] Proper descriptions

---

## 🧪 Testing Workflow

### Phase 1: Citizen Portal Testing

#### ✅ Test 1.1: Landing Page
**What to test:**
- [ ] Hero section displays correctly
- [ ] "Submit a Report" button is visible
- [ ] "Sign In" button is visible
- [ ] Three feature cards are displayed
- [ ] Responsive design works on mobile

**How to test:**
1. Visit the home page (`/`)
2. Verify all elements are visible
3. Test on different screen sizes

---

#### ✅ Test 1.2: User Registration
**What to test:**
- [ ] Sign up form accepts valid input
- [ ] All required fields are validated
- [ ] Profile is created successfully
- [ ] User is redirected to dashboard after signup

**How to test:**
1. Click "Sign In" on homepage
2. Switch to "Sign Up" tab
3. Fill in:
   - Full Name: "Test Citizen"
   - Phone: "+234 800 000 0001"
   - Email: `citizen@test.com`
   - Password: `test123456`
4. Submit and verify redirect to `/dashboard`

**Expected result:** 
- Success toast notification
- Redirect to citizen dashboard
- Profile created in database

---

#### ✅ Test 1.3: User Login
**What to test:**
- [ ] Login with correct credentials works
- [ ] Login with wrong credentials fails
- [ ] Error messages are clear
- [ ] Redirect works based on role

**How to test:**
1. Sign out if logged in
2. Go to `/auth`
3. Try logging in with wrong password (should fail)
4. Log in with correct credentials
5. Verify redirect to appropriate portal

**Expected result:**
- Successful login redirects to dashboard
- Failed login shows error message

---

#### ✅ Test 1.4: Submit a Report (Manual)
**What to test:**
- [ ] Form validates all required fields
- [ ] Category dropdown works
- [ ] Emergency checkbox works
- [ ] Report is created successfully
- [ ] Unique report number is generated
- [ ] Report appears in dashboard

**How to test:**
1. Navigate to `/submit-report` or click "New Report"
2. Fill in the form:
   - **Category**: "Road Infrastructure"
   - **Title**: "Pothole on Lagos-Ibadan Expressway"
   - **Description**: "Large pothole near KM 15, causing severe traffic. Approximately 3 feet wide and 8 inches deep. Multiple vehicles damaged."
   - **Location**: "Lagos-Ibadan Expressway, KM 15, Near Magboro"
   - **Emergency**: Leave unchecked
3. Submit the report
4. Go to dashboard and verify report appears

**Expected result:**
- Success toast with report number (e.g., `CL-2025-000001`)
- Redirect to dashboard
- Report shows with "SUBMITTED" status
- Report assigned to correct agency

---

#### ✅ Test 1.5: Submit Emergency Report
**What to test:**
- [ ] Emergency flag is saved correctly
- [ ] Emergency badge displays
- [ ] Higher priority handling

**How to test:**
1. Submit another report
2. This time, check the "Emergency" checkbox
3. Use:
   - **Category**: "Security"
   - **Title**: "Armed robbery in progress"
   - **Description**: "Suspected armed robbery at XYZ Bank branch"
   - **Location**: "XYZ Bank, Allen Avenue"
4. Verify emergency badge appears in dashboard

**Expected result:**
- Emergency badge visible
- Report marked as emergency in database

---

#### ✅ Test 1.6: View Reports Dashboard
**What to test:**
- [ ] All user's reports are displayed
- [ ] Reports show correct information
- [ ] Status badges display correctly
- [ ] Date formatting is correct
- [ ] Empty state shows when no reports

**How to test:**
1. Go to `/dashboard`
2. Verify both reports appear
3. Check each report card displays:
   - Report number
   - Title
   - Description
   - Location
   - Category
   - Status badge
   - Emergency badge (if applicable)
   - Creation date

**Expected result:**
- Both reports visible
- Information is accurate
- UI is clear and readable

---

### Phase 2: AI Classification Testing

#### ✅ Test 2.1: Automatic Report Classification
**What to test:**
- [ ] AI correctly identifies category
- [ ] AI assigns appropriate priority
- [ ] AI detects emergency situations
- [ ] Report is auto-assigned to agency

**How to test:**
1. Submit a report with ambiguous category:
   - **Category**: Select "Other" initially
   - **Title**: "Contaminated water supply"
   - **Description**: "The water from our taps has been brown and smells bad for 3 days. Multiple households affected. Children are getting sick."
2. Check if AI reclassifies to "Water & Sanitation"

**NOTE:** AI classification requires calling the edge function. This can be done manually or automatically on submission. Currently, you'll need to trigger it by:

```javascript
// In browser console after submitting:
await fetch(`${import.meta.env.VITE_SUPABASE_URL}/functions/v1/classify-report`, {
  method: 'POST',
  headers: {
    'Authorization': `Bearer ${import.meta.env.VITE_SUPABASE_PUBLISHABLE_KEY}`,
    'Content-Type': 'application/json'
  },
  body: JSON.stringify({
    reportId: 'YOUR_REPORT_ID',
    title: 'Report title',
    description: 'Report description'
  })
});
```

**Expected result:**
- Category updated to appropriate value
- Priority set based on severity
- Agency auto-assigned

---

### Phase 3: Officer Portal Testing

#### ✅ Test 3.1: Officer Access Control
**What to test:**
- [ ] Only officers can access `/officer`
- [ ] Citizens are redirected away
- [ ] Officer sees only their agency's reports

**How to test:**
1. Log in as citizen account
2. Try to access `/officer` directly
3. Verify you're redirected to dashboard
4. Log out and log in as officer account
5. Verify you're redirected to `/officer`

**Expected result:**
- Access control works correctly
- Officer portal loads successfully

---

#### ✅ Test 3.2: View Agency Reports
**What to test:**
- [ ] Officer sees reports assigned to their agency
- [ ] Reports from other agencies are hidden
- [ ] Statistics are accurate
- [ ] Filtering works

**How to test:**
1. As officer, view the main portal page
2. Check the stats cards (Total, Pending, In Progress, Resolved)
3. Verify only reports for your agency appear
4. Test the status filter dropdown

**Expected result:**
- Only relevant reports visible
- Stats match report counts
- Filter works correctly

---

#### ✅ Test 3.3: Update Report Status
**What to test:**
- [ ] Status can be updated
- [ ] Notes can be added
- [ ] Changes are saved
- [ ] Citizen can see updates

**How to test:**
1. Click "Update Status" on a report
2. Change status to "In Progress"
3. Add notes: "Team dispatched to location. ETA 2 hours."
4. Submit update
5. Log in as citizen and verify update appears

**Expected result:**
- Status updates successfully
- Notes are saved
- Citizen sees status change in dashboard

---

#### ✅ Test 3.4: Resolve Report
**What to test:**
- [ ] Report can be marked as resolved
- [ ] Resolved timestamp is set
- [ ] Report moves to resolved section

**How to test:**
1. Update a report to "Resolved" status
2. Add notes: "Pothole filled and road surface repaired."
3. Check if `resolved_at` timestamp is set in database

**Expected result:**
- Report status = "Resolved"
- Resolved count increases
- Timestamp recorded

---

### Phase 4: Admin Portal Testing

#### ✅ Test 4.1: Admin Access Control
**What to test:**
- [ ] Only admins can access `/admin`
- [ ] Non-admins are redirected
- [ ] Admin sees system-wide data

**How to test:**
1. Log in as citizen or officer
2. Try accessing `/admin`
3. Verify redirect
4. Log in as admin
5. Access should work

**Expected result:**
- Access control enforced
- Admin portal loads

---

#### ✅ Test 4.2: View System Statistics
**What to test:**
- [ ] Total reports count is accurate
- [ ] Total citizens count is correct
- [ ] Total officers count is correct
- [ ] Total agencies count is correct

**How to test:**
1. View admin dashboard
2. Compare stats with database counts

**Expected result:**
- All statistics match database

---

#### ✅ Test 4.3: Manage Agencies
**What to test:**
- [ ] All agencies are listed
- [ ] Agency details are displayed
- [ ] Can view agency information

**How to test:**
1. Go to "Agencies" tab
2. Verify all 10 agencies appear
3. Check each has:
   - Name
   - Category
   - Contact info
   - Active status

**Expected result:**
- All agencies visible
- Information is accurate

---

#### ✅ Test 4.4: View Officers
**What to test:**
- [ ] All officers are listed
- [ ] Officer details show correctly
- [ ] Agency assignments are visible

**How to test:**
1. Go to "Officers" tab
2. Verify officer list shows:
   - Officer name
   - Phone number
   - Assigned agency
   - Status

**Expected result:**
- Officers listed correctly
- Agency links work

---

### Phase 5: Authentication & Security Testing

#### ✅ Test 5.1: Row-Level Security (RLS)
**What to test:**
- [ ] Citizens can only see their own reports
- [ ] Officers can only see their agency's reports
- [ ] Admins can see all reports

**How to test:**
1. Create reports with Citizen A
2. Log in as Citizen B
3. Verify Citizen B cannot see Citizen A's reports
4. Log in as officer from same agency
5. Verify officer can see reports assigned to agency
6. Log in as admin
7. Verify admin can see all reports

**Expected result:**
- RLS policies enforce proper access control

---

#### ✅ Test 5.2: Profile Privacy
**What to test:**
- [ ] Users can only view/edit their own profile
- [ ] Officers can view profiles within scope
- [ ] Admins can view all profiles

**How to test:**
1. Try to access another user's profile
2. Verify access is denied
3. Test with each role level

**Expected result:**
- Privacy is maintained
- Proper access levels enforced

---

#### ✅ Test 5.3: Session Management
**What to test:**
- [ ] Session persists after page refresh
- [ ] Logout works correctly
- [ ] Login redirects appropriately

**How to test:**
1. Log in as any user
2. Refresh the page
3. Verify still logged in
4. Click logout
5. Verify redirected to home
6. Try accessing protected route
7. Verify redirected to login

**Expected result:**
- Session handling works correctly

---

### Phase 6: Edge Cases & Error Handling

#### ✅ Test 6.1: Form Validation
**What to test:**
- [ ] Required fields are enforced
- [ ] Email validation works
- [ ] Password requirements enforced
- [ ] Field length limits work

**How to test:**
1. Try submitting forms with:
   - Empty required fields
   - Invalid email format
   - Password too short
   - Description too long (>1000 chars)

**Expected result:**
- Appropriate error messages
- Form doesn't submit with invalid data

---

#### ✅ Test 6.2: Network Error Handling
**What to test:**
- [ ] Error messages display when API calls fail
- [ ] Loading states show during requests
- [ ] Graceful degradation

**How to test:**
1. Temporarily disconnect internet
2. Try submitting a report
3. Verify error message displays
4. Reconnect and retry

**Expected result:**
- Clear error messages
- No app crashes

---

#### ✅ Test 6.3: Empty States
**What to test:**
- [ ] Empty dashboard shows helpful message
- [ ] No reports state displays correctly
- [ ] Empty filters show appropriate message

**How to test:**
1. Create new user account
2. Check dashboard before submitting reports
3. Filter reports with no matches
4. Check officer portal with no assigned reports

**Expected result:**
- Helpful empty state messages
- Call-to-action buttons present

---

## 📊 Database Verification

### Key Tables to Check

1. **profiles**
   - Verify user roles are set correctly
   - Check agency assignments for officers

2. **reports**
   - Verify report numbers are unique
   - Check agency assignments
   - Verify timestamps

3. **report_updates**
   - Check update notes are saved
   - Verify updated_by references

4. **agencies**
   - All 10 agencies present
   - Categories assigned correctly

5. **audit_logs**
   - Actions are being logged
   - Timestamps are recorded

---

## 🔍 Performance Testing

### Page Load Times
- [ ] Home page loads < 2 seconds
- [ ] Dashboard loads < 3 seconds
- [ ] Report submission completes < 2 seconds

### Database Query Performance
- [ ] Report listing is fast (< 1 second)
- [ ] Filtering doesn't cause lag
- [ ] Large report lists paginate properly

---

## 📱 Responsive Design Testing

Test on these viewport sizes:
- [ ] Mobile (375px width)
- [ ] Tablet (768px width)
- [ ] Desktop (1440px width)

Check:
- [ ] All buttons are tappable
- [ ] Text is readable
- [ ] Forms are usable
- [ ] Navigation works
- [ ] Cards stack properly

---

## 🐛 Known Limitations (Not Yet Implemented)

These features are not included in the MVP:

1. **WhatsApp Integration**: No chatbot yet
2. **Media Upload**: Cannot upload photos/videos
3. **Face/License Plate Blur**: Image processing not implemented
4. **NIN Verification**: Mock structure only, no real API
5. **Email Notifications**: Not configured
6. **SMS Notifications**: Not configured
7. **Real-time Updates**: No websocket connections
8. **Export to CSV**: Not implemented
9. **Advanced Analytics**: Basic stats only
10. **Multi-language Support**: English only

---

## ✅ Testing Completion Checklist

### Core Features
- [ ] User registration works
- [ ] User login works
- [ ] Citizens can submit reports
- [ ] Reports appear in dashboard
- [ ] Officers can view agency reports
- [ ] Officers can update report status
- [ ] Admins can view all data
- [ ] RLS security works

### AI Features
- [ ] AI classification endpoint exists
- [ ] Classification updates reports
- [ ] Agency assignment works

### UI/UX
- [ ] Design is professional
- [ ] Responsive on all devices
- [ ] Forms validate correctly
- [ ] Error messages are clear
- [ ] Loading states are visible

### Security
- [ ] Access control enforced
- [ ] Data privacy maintained
- [ ] Passwords are secure
- [ ] Sessions work correctly

---

## 🎓 Tips for Effective Testing

1. **Use Different Browsers**: Test on Chrome, Firefox, Safari
2. **Clear Cache**: Between major changes
3. **Test Real Scenarios**: Think like actual users
4. **Document Issues**: Take screenshots of bugs
5. **Test Edge Cases**: Try to break things
6. **Verify Database**: Check data is being saved correctly

---

## 📝 Bug Reporting Template

When you find issues, report them like this:

```
**Bug Title**: [Clear description]

**Steps to Reproduce**:
1. [Step 1]
2. [Step 2]
3. [Step 3]

**Expected Result**: [What should happen]
**Actual Result**: [What actually happened]
**User Role**: [Citizen/Officer/Admin]
**Browser**: [Chrome/Firefox/Safari]
**Screenshots**: [If applicable]
```

---

## 🚀 Next Steps After Testing

Once testing is complete:

1. **Document Findings**: Create a summary of all tests
2. **Prioritize Bugs**: Critical, High, Medium, Low
3. **Plan Iterations**: Decide what to build next
4. **User Acceptance**: Have real users test
5. **Performance Optimization**: If needed
6. **Deploy to Production**: When ready

---

## 📞 Support

If you encounter issues during testing:
- Check Lovable Cloud logs (View Backend → Functions)
- Review browser console for errors
- Verify database state
- Check authentication status

---

**Testing Status**: [ ] Not Started | [ ] In Progress | [ ] Completed

**Last Updated**: [Date]
**Tester Name**: [Your Name]
