package tableau

import (
	"fmt"
	"github.com/tiketdatarisal/tableau/models"
	"net/http"
	. "net/url"
)

type workbooksViews struct {
	base *Client
}

// AddTagsToView Adds one or more tags to the specified view.
//
// URI:
//   PUT /api/api-version/sites/site-id/views/view-id/tags
// Reference: https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_workbooks_and_views.htm#add_tags_to_view
func (w *workbooksViews) AddTagsToView(viewID string, tagNames []string) ([]models.Tag, error) {
	if !w.base.Authentication.IsSignedIn() {
		if err := w.base.Authentication.SignIn(); err != nil {
			return nil, err
		}
	}

	var tags []models.Tag
	for _, tagName := range tagNames {
		if tagName != "" {
			tags = append(tags, models.Tag{Label: tagName})
		}
	}

	if len(tags) == 0 {
		return nil, ErrBadRequest
	}

	reqBody := models.TagBody{
		Tags: (*struct {
			Tag []models.Tag `json:"tag,omitempty"`
		})(&struct{ Tag []models.Tag }{Tag: tags}),
	}

	url := w.base.cfg.GetUrl(fmt.Sprintf(addTagsToViewPath, w.base.Authentication.siteID, viewID))
	if url == "" {
		return nil, ErrInvalidHost
	}

	res, err := w.base.c.R().
		SetHeader(contentTypeHeader, mimeTypeJSON).
		SetHeader(acceptHeader, mimeTypeJSON).
		SetHeader(authorizationHeader, w.base.Authentication.getBearerToken()).
		SetBody(reqBody).
		Put(url)

	if err != nil {
		errBody, err := models.NewErrorBody(res.Body())
		if err != nil {
			return nil, ErrUnknownError
		}

		return nil, errCodeMap[errBody.Error.Code]
	}

	if res.StatusCode() != http.StatusOK {
		errBody, err := models.NewErrorBody(res.Body())
		if err != nil {
			return nil, ErrUnknownError
		}

		return nil, errCodeMap[errBody.Error.Code]
	}

	resBody := models.TagBody{}
	if err = json.Unmarshal(res.Body(), &resBody); err != nil {
		return nil, ErrFailedUnmarshalResponseBody
	}

	if resBody.Tags == nil {
		return nil, nil
	}

	return resBody.Tags.Tag, nil
}

// AddTagsToWorkbook Adds one or more tags to the specified workbook.
//
// URI:
//   PUT /api/api-version/sites/site-id/workbooks/workbook-id/tags
// Reference: https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_workbooks_and_views.htm#add_tags_to_workbook
func (w *workbooksViews) AddTagsToWorkbook(workbookID string, tagNames []string) ([]models.Tag, error) {
	if !w.base.Authentication.IsSignedIn() {
		if err := w.base.Authentication.SignIn(); err != nil {
			return nil, err
		}
	}

	var tags []models.Tag
	for _, tagName := range tagNames {
		if tagName != "" {
			tags = append(tags, models.Tag{Label: tagName})
		}
	}

	if len(tags) == 0 {
		return nil, ErrBadRequest
	}

	reqBody := models.TagBody{
		Tags: (*struct {
			Tag []models.Tag `json:"tag,omitempty"`
		})(&struct{ Tag []models.Tag }{Tag: tags}),
	}

	url := w.base.cfg.GetUrl(fmt.Sprintf(addTagsToWorkbookPath, w.base.Authentication.siteID, workbookID))
	if url == "" {
		return nil, ErrInvalidHost
	}

	res, err := w.base.c.R().
		SetHeader(contentTypeHeader, mimeTypeJSON).
		SetHeader(acceptHeader, mimeTypeJSON).
		SetHeader(authorizationHeader, w.base.Authentication.getBearerToken()).
		SetBody(reqBody).
		Put(url)

	if err != nil {
		errBody, err := models.NewErrorBody(res.Body())
		if err != nil {
			return nil, ErrUnknownError
		}

		return nil, errCodeMap[errBody.Error.Code]
	}

	if res.StatusCode() != http.StatusOK {
		errBody, err := models.NewErrorBody(res.Body())
		if err != nil {
			return nil, ErrUnknownError
		}

		return nil, errCodeMap[errBody.Error.Code]
	}

	resBody := models.TagBody{}
	if err = json.Unmarshal(res.Body(), &resBody); err != nil {
		return nil, ErrFailedUnmarshalResponseBody
	}

	if resBody.Tags == nil {
		return nil, nil
	}

	return resBody.Tags.Tag, nil
}

// DeleteTagFromView Deletes a tag from the specified view.
//
// URI:
//   DELETE /api/api-version/sites/site-id/views/view-id/tags/tag-name
// Reference: https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_workbooks_and_views.htm#delete_tag_from_view
func (w *workbooksViews) DeleteTagFromView(viewID, tagName string) error {
	if !w.base.Authentication.IsSignedIn() {
		if err := w.base.Authentication.SignIn(); err != nil {
			return err
		}
	}

	url := w.base.cfg.GetUrl(fmt.Sprintf(deleteTagFromViewPath, w.base.Authentication.siteID, viewID, QueryEscape(tagName)))
	if url == "" {
		return ErrInvalidHost
	}

	res, err := w.base.c.R().
		SetHeader(contentTypeHeader, mimeTypeJSON).
		SetHeader(acceptHeader, mimeTypeJSON).
		SetHeader(authorizationHeader, w.base.Authentication.getBearerToken()).
		Delete(url)
	if err != nil {
		errBody, err := models.NewErrorBody(res.Body())
		if err != nil {
			return ErrUnknownError
		}

		return errCodeMap[errBody.Error.Code]
	}

	if res.StatusCode() != http.StatusNoContent {
		errBody, err := models.NewErrorBody(res.Body())
		if err != nil {
			return ErrUnknownError
		}

		return errCodeMap[errBody.Error.Code]
	}

	return nil
}

// DeleteTagFromWorkbook Deletes a tag from the specified workbook.
//
// URI:
//   DELETE /api/api-version/sites/site-id/workbooks/workbook-id/tags/tag-name
// Reference: https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_workbooks_and_views.htm#delete_tag_from_workbook
func (w *workbooksViews) DeleteTagFromWorkbook(workbookID, tagName string) error {
	if !w.base.Authentication.IsSignedIn() {
		if err := w.base.Authentication.SignIn(); err != nil {
			return err
		}
	}

	url := w.base.cfg.GetUrl(fmt.Sprintf(deleteTagFromWorkbookPath, w.base.Authentication.siteID, workbookID, QueryEscape(tagName)))
	if url == "" {
		return ErrInvalidHost
	}

	res, err := w.base.c.R().
		SetHeader(contentTypeHeader, mimeTypeJSON).
		SetHeader(acceptHeader, mimeTypeJSON).
		SetHeader(authorizationHeader, w.base.Authentication.getBearerToken()).
		Delete(url)
	if err != nil {
		errBody, err := models.NewErrorBody(res.Body())
		if err != nil {
			return ErrUnknownError
		}

		return errCodeMap[errBody.Error.Code]
	}

	if res.StatusCode() != http.StatusNoContent {
		errBody, err := models.NewErrorBody(res.Body())
		if err != nil {
			return ErrUnknownError
		}

		return errCodeMap[errBody.Error.Code]
	}

	return nil
}

// DownloadWorkbookPDF Downloads a .pdf containing images of the sheets that the user has permission to view in a workbook.
// Download Images/PDF permissions must be enabled for the workbook (true by default).
// If Show sheets in tabs is not selected for the workbook, only the default tab will appear in the .pdf file.
//
// If you make multiple requests for a PDF, subsequent calls return a cached version of the file.
// This means that the returned PDF might not include the latest changes to the workbook.
// To decrease the amount of time that a workbook is cached, use the maxAge parameter.
//
// URI:
//   GET /api/api-version/sites/site-id/workbooks/workbook-id/pdf?type=page-type&orientation=page-orientation
// Reference: https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_workbooks_and_views.htm#download_workbook_pdf
func (w *workbooksViews) DownloadWorkbookPDF() {}

// GetView Gets the details of a specific view.
//
// URI:
//   GET /api/api-version/sites/site-id/views/view-id
// Reference: https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_workbooks_and_views.htm#get_view
func (w *workbooksViews) GetView() {}

// GetViewByPath Gets the details of all views in a site with a specified name.
//
// URI:
//   GET /api/api-version/sites/site-id/views?filter=viewUrlName:eq:view-name
// Reference: https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_workbooks_and_views.htm#get_view_by_path
func (w *workbooksViews) GetViewByPath() {}

// QueryViewsForSite Returns all the views for the specified site, optionally including usage statistics.
//
// URI:
//   GET /api/api-version/sites/site-id/views
// Reference: https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_workbooks_and_views.htm#query_views_for_site
func (w *workbooksViews) QueryViewsForSite() {}

// QueryViewsForWorkbook Returns all the views for the specified workbook, optionally including usage statistics.
//
// URI:
//   GET /api/api-version/sites/site-id/workbooks/workbook-id/views
// Reference: https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_workbooks_and_views.htm#query_views_for_workbook
func (w *workbooksViews) QueryViewsForWorkbook() {}

// QueryViewImage Returns an image of the specified view.
// If you make multiple requests for an image, subsequent calls return a cached version of the image.
// This means that the returned image might not include the latest changes to the view.
// To decrease the amount of time that an image is cached, use the maxAge parameter.
//
// URI:
//   GET /api/api-version/sites/site-id/views/view-id/image
// Reference: https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_workbooks_and_views.htm#query_view_image
func (w *workbooksViews) QueryViewImage(viewID string, maxAgeInMinutes ...int) ([]byte, error) {
	if !w.base.Authentication.IsSignedIn() {
		if err := w.base.Authentication.SignIn(); err != nil {
			return nil, err
		}
	}

	maxAge := defaultMaxAge
	if len(maxAgeInMinutes) > 0 {
		maxAge = 1
		if maxAgeInMinutes[0] > 1 {
			maxAge = maxAgeInMinutes[0]
		}
	}

	url := w.base.cfg.GetUrl(fmt.Sprintf(queryViewImagePath, w.base.Authentication.siteID, viewID))
	if url == "" {
		return nil, ErrInvalidHost
	}

	url = fmt.Sprintf(viewImageParams, url, maxAge)

	res, err := w.base.c.R().
		SetHeader(contentTypeHeader, mimeTypeJSON).
		SetHeader(acceptHeader, mimeTypeAny).
		SetHeader(authorizationHeader, w.base.Authentication.getBearerToken()).
		Get(url)

	if err != nil {
		errBody, err := models.NewErrorBody(res.Body())
		if err != nil {
			return nil, ErrUnknownError
		}

		return nil, errCodeMap[errBody.Error.Code]
	}

	if res.StatusCode() != http.StatusOK {
		errBody, err := models.NewErrorBody(res.Body())
		if err != nil {
			return nil, ErrUnknownError
		}

		return nil, errCodeMap[errBody.Error.Code]
	}

	return res.Body(), nil
}

// QueryViewPDF Returns a specified view rendered as a .pdf file.
// If you make multiple requests for a PDF, subsequent calls return a cached version of the file.
// This means that the returned PDF might not include the latest changes to the view.
// To decrease the amount of time that an PDF is cached, use the maxAge parameter.
//
// URI:
//   GET /api/api-version/sites/site-id/views/view-id/pdf
// Reference: https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_workbooks_and_views.htm#query_view_pdf
func (w *workbooksViews) QueryViewPDF() {}

// QueryWorkbook Returns information about the specified workbook, including information about views and tags.
//
// URI:
//   GET /api/api-version/sites/site-id/workbooks/workbook-id
// Reference: https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_workbooks_and_views.htm#query_workbook
func (w *workbooksViews) QueryWorkbook() {}

// QueryWorkbooksForSite Returns the workbooks on a site.
// If the user is not an administrator, the method returns just the workbooks that the user has permissions to view.
//
// URI:
//   GET /api/api-version/sites/site-id/workbooks
// Reference: https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_workbooks_and_views.htm#query_workbooks_for_site
func (w *workbooksViews) QueryWorkbooksForSite() {}

// QueryWorkbooksForUser Returns the workbooks that the specified user owns in addition to those that the user has Read (view) permissions for.
//
// URI:
//   GET /api/api-version/sites/site-id/users/user-id/workbooks
// Reference: https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_workbooks_and_views.htm#query_workbooks_for_user
func (w *workbooksViews) QueryWorkbooksForUser() {}